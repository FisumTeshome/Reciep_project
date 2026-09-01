/**
 * Sanitizes API and runtime errors, extracting backend messages or mapping HTTP status codes,
 * while stripping out technical FetchErrors, URLs, ports, and raw endpoints.
 */
export function parseApiError(err: any, fallbackMessage: string = 'An error occurred. Please try again.'): string {
  if (!err) return fallbackMessage;

  // If err is a plain string
  if (typeof err === 'string') {
    if (err.includes('FetchError:') || err.includes('http://') || err.includes('https://') || err.includes('localhost')) {
      return fallbackMessage;
    }
    return err;
  }

  // Extract backend JSON error object if returned
  const data = err.data || err.response?._data || err.response?.data || err.value?.data;
  if (data) {
    if (typeof data.error === 'string' && data.error.trim()) {
      return data.error;
    }
    if (typeof data.message === 'string' && data.message.trim()) {
      return data.message;
    }
    if (typeof data.detail === 'string' && data.detail.trim()) {
      return data.detail;
    }
  }

  // Handle specific HTTP status codes
  const status = err.status || err.statusCode || err.response?.status || err.value?.status || err.value?.statusCode;
  if (status === 401) {
    return 'Invalid credentials or session expired. Please log in again.';
  }
  if (status === 403) {
    return 'You do not have permission to perform this action.';
  }
  if (status === 404) {
    return 'The requested resource was not found.';
  }
  if (status === 409) {
    return 'A record with this information already exists.';
  }
  if (status >= 500) {
    return 'Server error. Please try again later.';
  }

  // Fallback to err.message if string does not expose FetchError or URLs
  if (typeof err.message === 'string' && err.message.trim()) {
    const msg = err.message.trim();
    if (!msg.includes('FetchError:') && !msg.includes('http://') && !msg.includes('https://') && !msg.includes('localhost')) {
      return msg;
    }
  }

  return fallbackMessage;
}
