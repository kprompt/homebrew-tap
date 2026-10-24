/**
 * Fetcher for the latest default models.
 * 
 * This module simulates fetching the latest default models from a remote source.
 */

import { DefaultModel } from './defaultModelRefresher';

/**
 * Simulates fetching the latest default models.
 */
export async function fetchLatestDefaultModels(): Promise<Omit<DefaultModel, 'lastUpdated'>[]> {
  // Simulate network delay
  await new Promise(resolve => setTimeout(resolve, 50));

  // Return mocked latest models
  return [
    { id: 'modelA', data: { name: 'Model A', version: '1.1' } },
    { id: 'modelB', data: { name: 'Model B', version: '2.0' } },
  ];
}
