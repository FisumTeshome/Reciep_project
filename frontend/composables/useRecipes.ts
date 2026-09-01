import { useRuntimeConfig, useState, useFetch } from '#imports';
import { parseApiError } from '~/utils/errorUtils';

export function useRecipes() {
  const config = useRuntimeConfig();
  const apiBase = config.public.apiBase;
  const recipes = useState<any[]>('recipeList', () => []);
  const recipe = useState<any | null>('currentRecipe', () => null);
  const loading = useState<boolean>('recipesLoading', () => false);
  const error = useState<string | null>('recipesError', () => null);

  const fetchRecipes = async (params: { search?: string; category_id?: string } = {}) => {
    loading.value = true;
    error.value = null;

    const query = new URLSearchParams();
    if (params.search) query.append('search', params.search);
    if (params.category_id) query.append('category_id', params.category_id);

    try {
      const { data, error: fetchError } = await useFetch(`${apiBase}/api/recipes?${query.toString()}`, {
        method: 'GET',
      });

      if (fetchError.value) {
        error.value = parseApiError(fetchError.value, 'Failed to load recipes. Please try again later.');
        recipes.value = [];
      } else {
        recipes.value = (data.value as any[]) || [];
      }
    } catch (err) {
      error.value = parseApiError(err, 'Failed to load recipes. Please try again later.');
      recipes.value = [];
    } finally {
      loading.value = false;
    }

    return recipes.value;
  };

  const fetchRecipe = async (id: string) => {
    loading.value = true;
    error.value = null;

    try {
      const { data, error: fetchError } = await useFetch(`${apiBase}/api/recipes/get?id=${encodeURIComponent(id)}`, {
        method: 'GET',
      });

      if (fetchError.value) {
        error.value = parseApiError(fetchError.value, 'Failed to load recipe details.');
        recipe.value = null;
      } else {
        recipe.value = data.value || null;
      }
    } catch (err) {
      error.value = parseApiError(err, 'Failed to load recipe details.');
      recipe.value = null;
    } finally {
      loading.value = false;
    }

    return recipe.value;
  };

  return {
    apiBase,
    recipes,
    recipe,
    loading,
    error,
    fetchRecipes,
    fetchRecipe,
  };
}
