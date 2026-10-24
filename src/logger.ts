/**
 * Simple logger utility.
 * Replace with real logger in production.
 */

export const logger = {
  info: (...args: any[]) => {
    console.info('[INFO]', ...args);
  },
  error: (...args: any[]) => {
    console.error('[ERROR]', ...args);
  },
};
