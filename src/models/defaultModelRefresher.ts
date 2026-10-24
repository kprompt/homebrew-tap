/**
 * Module to refresh stale default models.
 * 
 * This module exports a function `refreshDefaultModels` that checks if the default models
 * are stale based on a configurable TTL and refreshes them if needed.
 */

import { getDefaultModels, saveDefaultModels } from './modelStorage';
import { fetchLatestDefaultModels } from './modelFetcher';

const DEFAULT_MODEL_TTL_MS = 24 * 60 * 60 * 1000; // 24 hours

interface DefaultModel {
  id: string;
  data: any;
  lastUpdated: number;
}

/**
 * Checks if the default models are stale.
 * @param models - Array of default models.
 * @returns true if any model is stale, false otherwise.
 */
function areModelsStale(models: DefaultModel[]): boolean {
  const now = Date.now();
  return models.some(model => (now - model.lastUpdated) > DEFAULT_MODEL_TTL_MS);
}

/**
 * Refreshes stale default models.
 * - Loads current default models.
 * - If stale, fetches latest models and saves them.
 * - Returns the refreshed models.
 */
export async function refreshDefaultModels(): Promise<DefaultModel[]> {
  const currentModels = await getDefaultModels();

  if (!currentModels || currentModels.length === 0 || areModelsStale(currentModels)) {
    const latestModels = await fetchLatestDefaultModels();
    const now = Date.now();
    const updatedModels = latestModels.map(model => ({
      ...model,
      lastUpdated: now,
    }));
    await saveDefaultModels(updatedModels);
    return updatedModels;
  }

  return currentModels;
}
