<script setup lang="ts">
import { useRouter, useRoute } from '#imports';
import { useRecipes } from '~/composables/useRecipes';
import { useAuth } from '~/composables/useAuth';
import { useRuntimeConfig } from '#imports';
import { parseApiError } from '~/utils/errorUtils';

const router = useRouter();
const route = useRoute();
const auth = useAuth();
const { recipe, loading, fetchRecipe } = useRecipes();
const config = useRuntimeConfig();
const apiBase = config.public.apiBase;

const formLoading = ref(false);
const error = ref('');
const recipeId = route.params.id as string;

// Recipe form data
const title = ref('');
const description = ref('');
const category = ref('');
const prepTime = ref('');
const cookTime = ref('');
const servings = ref('');
const difficulty = ref('Medium');
const featuredImage = ref('');

// Ingredients
const ingredients = ref([{ name: '', quantity: '', unit: '' }]);

// Steps
const steps = ref([{ description: '', order: 1 }]);

// Load existing recipe data
onMounted(async () => {
  try {
    await fetchRecipe(recipeId);
    if (recipe.value) {
      title.value = recipe.value.title || '';
      description.value = recipe.value.description || '';
      category.value = recipe.value.category?.name || recipe.value.category || '';
      prepTime.value = recipe.value.prep_time?.toString() || '';
      cookTime.value = recipe.value.cook_time?.toString() || '';
      servings.value = recipe.value.servings?.toString() || '';
      difficulty.value = recipe.value.difficulty || 'Medium';
      featuredImage.value = recipe.value.featured_image || '';
      
      // Load ingredients
      if (recipe.value.recipe_ingredients && recipe.value.recipe_ingredients.length > 0) {
        ingredients.value = recipe.value.recipe_ingredients.map((ing: any) => ({
          name: ing.ingredient?.name || ing.name || '',
          quantity: ing.quantity || '',
          unit: ing.unit || ''
        }));
      }
      
      // Load steps
      if (recipe.value.recipe_steps && recipe.value.recipe_steps.length > 0) {
        steps.value = recipe.value.recipe_steps.map((step: any) => ({
          description: step.description || '',
          order: step.order || step.step_number || 1
        })).sort((a: any, b: any) => a.order - b.order);
      }
    }
  } catch (err) {
    error.value = parseApiError(err, 'Failed to load recipe');
  }
});

const addIngredient = () => {
  ingredients.value.push({ name: '', quantity: '', unit: '' });
};

const removeIngredient = (index: number) => {
  if (ingredients.value.length > 1) {
    ingredients.value.splice(index, 1);
  }
};

const addStep = () => {
  steps.value.push({ description: '', order: steps.value.length + 1 });
};

const removeStep = (index: number) => {
  if (steps.value.length > 1) {
    steps.value.splice(index, 1);
    steps.value.forEach((step, i) => {
      step.order = i + 1;
    });
  }
};

const submit = async () => {
  error.value = '';
  formLoading.value = true;

  try {
    const recipeData = {
      title: title.value,
      description: description.value,
      category: category.value,
      prep_time: parseInt(prepTime.value) || 0,
      cook_time: parseInt(cookTime.value) || 0,
      servings: parseInt(servings.value) || 0,
      difficulty: difficulty.value,
      featured_image: featuredImage.value,
      ingredients: ingredients.value.filter(ing => ing.name.trim() !== ''),
      steps: steps.value.filter(step => step.description.trim() !== ''),
    };

    const { data, error: fetchError } = await useFetch(`${apiBase}/api/recipes/update?id=${encodeURIComponent(recipeId)}`, {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${auth.token.value}`,
        'Content-Type': 'application/json',
      },
      body: recipeData,
    });

    if (fetchError.value) {
      throw fetchError.value;
    }

    await router.push(`/recipes/${recipeId}`);
  } catch (err: any) {
    error.value = parseApiError(err, 'Failed to update recipe');
  } finally {
    formLoading.value = false;
  }
};
</script>

<template>
  <div class="space-y-8 px-4 py-12 sm:px-6 lg:px-8">
    <div v-if="loading" class="rounded-[32px] border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-800 dark:bg-slate-900">
      Loading recipe...
    </div>

    <div v-else class="rounded-[32px] border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-800 dark:bg-slate-900">
      <h1 class="text-3xl font-black text-slate-900 dark:text-white">Edit Recipe</h1>
      <p class="mt-2 text-sm text-slate-500 dark:text-slate-400">Update your recipe details.</p>

      <div v-if="error" class="mt-6 rounded-2xl bg-red-50 px-4 py-3 text-sm text-red-700 dark:bg-red-900/20 dark:text-red-200">
        {{ error }}
      </div>

      <form @submit.prevent="submit" class="mt-8 space-y-8">
        <!-- Basic Information -->
        <section class="space-y-6">
          <h2 class="text-xl font-bold text-slate-900 dark:text-white">Basic Information</h2>
          
          <div>
            <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Recipe Title</label>
            <input v-model="title" type="text" placeholder="e.g., Doro Wot" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
          </div>

          <div>
            <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Description</label>
            <textarea v-model="description" rows="3" placeholder="Describe your recipe..." required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white"></textarea>
          </div>

          <div class="grid gap-6 md:grid-cols-2">
            <div>
              <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Category</label>
              <select v-model="category" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white">
                <option value="">Select category</option>
                <option>Ethiopian</option>
                <option>African</option>
                <option>Desserts</option>
                <option>Vegetarian</option>
                <option>Breakfast</option>
                <option>Dinner</option>
              </select>
            </div>

            <div>
              <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Difficulty</label>
              <select v-model="difficulty" class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white">
                <option>Easy</option>
                <option>Medium</option>
                <option>Hard</option>
              </select>
            </div>
          </div>

          <div class="grid gap-6 md:grid-cols-3">
            <div>
              <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Prep Time (minutes)</label>
              <input v-model="prepTime" type="number" placeholder="30" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
            </div>

            <div>
              <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Cook Time (minutes)</label>
              <input v-model="cookTime" type="number" placeholder="45" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
            </div>

            <div>
              <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Servings</label>
              <input v-model="servings" type="number" placeholder="4" required class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
            </div>
          </div>

          <div>
            <label class="mb-2 block text-sm font-semibold text-slate-700 dark:text-slate-200">Featured Image URL</label>
            <input v-model="featuredImage" type="url" placeholder="https://example.com/image.jpg" class="w-full rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
          </div>
        </section>

        <!-- Ingredients -->
        <section class="space-y-6">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-bold text-slate-900 dark:text-white">Ingredients</h2>
            <button @click.prevent="addIngredient" type="button" class="rounded-full bg-orange-500 px-4 py-2 text-sm font-semibold text-white transition hover:bg-orange-600">
              + Add Ingredient
            </button>
          </div>

          <div class="space-y-4">
            <div v-for="(ingredient, index) in ingredients" :key="index" class="flex gap-4">
              <input v-model="ingredient.name" type="text" placeholder="Ingredient name" class="flex-1 rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
              <input v-model="ingredient.quantity" type="text" placeholder="Quantity" class="w-24 rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
              <input v-model="ingredient.unit" type="text" placeholder="Unit" class="w-24 rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white" />
              <button @click.prevent="removeIngredient(index)" type="button" class="rounded-full border border-red-300 bg-red-50 px-3 py-2 text-sm font-semibold text-red-600 transition hover:bg-red-100 dark:border-red-700 dark:bg-red-900/20 dark:text-red-400" :disabled="ingredients.length === 1">
                ✕
              </button>
            </div>
          </div>
        </section>

        <!-- Steps -->
        <section class="space-y-6">
          <div class="flex items-center justify-between">
            <h2 class="text-xl font-bold text-slate-900 dark:text-white">Instructions</h2>
            <button @click.prevent="addStep" type="button" class="rounded-full bg-orange-500 px-4 py-2 text-sm font-semibold text-white transition hover:bg-orange-600">
              + Add Step
            </button>
          </div>

          <div class="space-y-4">
            <div v-for="(step, index) in steps" :key="index" class="flex gap-4">
              <div class="flex h-12 w-12 items-center justify-center rounded-full bg-orange-500 text-white font-bold">
                {{ step.order }}
              </div>
              <textarea v-model="step.description" rows="2" :placeholder="`Step ${step.order} instructions...`" class="flex-1 rounded-2xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm text-slate-900 outline-none transition focus:border-orange-400 focus:ring-2 focus:ring-orange-200 dark:border-slate-700 dark:bg-slate-800 dark:text-white"></textarea>
              <button @click.prevent="removeStep(index)" type="button" class="mt-2 rounded-full border border-red-300 bg-red-50 px-3 py-2 text-sm font-semibold text-red-600 transition hover:bg-red-100 dark:border-red-700 dark:bg-red-900/20 dark:text-red-400" :disabled="steps.length === 1">
                ✕
              </button>
            </div>
          </div>
        </section>

        <!-- Submit Button -->
        <div class="flex gap-4">
          <button type="button" @click="router.back" class="rounded-full border border-slate-200 bg-slate-100 px-6 py-3 text-sm font-semibold text-slate-700 transition hover:bg-slate-200 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700">
            Cancel
          </button>
          <button type="submit" :disabled="formLoading" class="flex-1 rounded-full bg-orange-500 px-6 py-3 text-sm font-semibold text-white transition hover:bg-orange-600 disabled:cursor-not-allowed disabled:opacity-60">
            {{ formLoading ? 'Updating Recipe...' : 'Update Recipe' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
