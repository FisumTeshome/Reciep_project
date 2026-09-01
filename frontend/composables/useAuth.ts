import { useRuntimeConfig, useState, computed, useFetch } from '#imports';
import { parseApiError } from '~/utils/errorUtils';

interface User {
  id: string;
  username: string;
  email: string;
  first_name?: string;
  last_name?: string;
  avatar?: string;
  bio?: string;
}

interface LoginPayload {
  identifier: string;
  password: string;
}

interface RegisterPayload {
  username: string;
  email: string;
  password: string;
}

export function useAuth() {
  const config = useRuntimeConfig();
  const apiBase = config.public.apiBase;
  const token = useState<string | null>('authToken', () => null);
  const user = useState<User | null>('authUser', () => null);
  const isInitialized = useState<boolean>('authInitialized', () => false);
  const isLoggedIn = computed(() => !!token.value && !!user.value);

  const loadToken = () => {
    if (process.client) {
      const stored = localStorage.getItem('recipejoy_token');
      token.value = stored || null;
      isInitialized.value = true;
    }
  };

  const saveToken = (value: string | null) => {
    token.value = value;
    if (process.client) {
      if (value) {
        localStorage.setItem('recipejoy_token', value);
      } else {
        localStorage.removeItem('recipejoy_token');
      }
    }
  };

  const fetchMe = async () => {
    if (!token.value) {
      user.value = null;
      return null;
    }
    const { data, error } = await useFetch(`${apiBase}/api/auth/me`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });

    if (error.value) {
      saveToken(null);
      user.value = null;
      throw new Error(parseApiError(error.value, 'Session expired or invalid token. Please log in again.'));
    }

    user.value = data.value as User;
    return user.value;
  };

  const login = async (payload: LoginPayload) => {
    const { data, error } = await useFetch(`${apiBase}/api/auth/login`, {
      method: 'POST',
      body: payload,
    });
    if (error.value || !data.value) {
      throw new Error(parseApiError(error.value, 'Invalid username/email or password.'));
    }
    const body = data.value as { token: string; user: User };
    saveToken(body.token);
    user.value = body.user;
    return body.user;
  };

  const register = async (payload: RegisterPayload) => {
    const { data, error } = await useFetch(`${apiBase}/api/auth/register`, {
      method: 'POST',
      body: payload,
    });
    if (error.value || !data.value) {
      throw new Error(parseApiError(error.value, 'Registration failed. Please check your information.'));
    }
    const body = data.value as { token: string; user: User };
    saveToken(body.token);
    user.value = body.user;
    return body.user;
  };

  const logout = async () => {
    const currentToken = token.value;
    saveToken(null);
    user.value = null;
    if (currentToken) {
      await $fetch(`${apiBase}/api/auth/logout`, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${currentToken}`,
        },
      }).catch(() => null);
    }
  };

  const initializeAuth = async () => {
    if (process.client && !isInitialized.value) {
      loadToken();
      if (token.value) {
        try {
          await fetchMe();
        } catch (err) {
          saveToken(null);
          user.value = null;
        }
      } else {
        user.value = null;
      }
    }
  };

  if (process.client && !isInitialized.value) {
    initializeAuth();
  }

  return {
    apiBase,
    token,
    user,
    isLoggedIn,
    fetchMe,
    login,
    register,
    logout,
    initializeAuth,
  };
}
