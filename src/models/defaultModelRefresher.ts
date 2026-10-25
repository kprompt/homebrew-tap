/**
 * Module to refresh stale default models.
 * 
 * This module exports a function `refreshStaleDefaultModels` which:
 * - Detects default models that are stale (older than a configured threshold)
 * - Refreshes them by reloading or rebuilding the model
 * - Returns a summary of refreshed models
 */

import { getDefaultModels, reloadModel, getModelLastUpdated } from './modelStorage';
import { Logger } from './logger';

const STALE_THRESHOLD_MS = 1000 * 60 * 60 * 24 * 7; // 7 days

export interface RefreshResult {
  refreshedModels: string[];
  skippedModels: string[];
}

/**
 * Refresh stale default models.
 * 
 * A default model is considered stale if its last updated timestamp is older than 7 days.
 * This function reloads all stale default models and returns a summary.
 */
export async function refreshStaleDefaultModels(logger: Logger): Promise<RefreshResult> {
  const defaultModels = await getDefaultModels();
  const now = Date.now();

  const refreshedModels: string[] = [];
  const skippedModels: string[] = [];

  for (const model of defaultModels) {
    const lastUpdated = await getModelLastUpdated(model.id);
    if (!lastUpdated) {
      logger.warn(`Model ${model.id} has no last updated timestamp, refreshing by default.`);
      await reloadModel(model.id);
      refreshedModels.push(model.id);
      continue;
    }
    if (now - lastUpdated > STALE_THRESHOLD_MS) {
      logger.info(`Model ${model.id} is stale, refreshing.`);
      await reloadModel(model.id);
      refreshedModels.push(model.id);
    } else {
      skippedModels.push(model.id);
    }
  }

  logger.info(`Refreshed ${refreshedModels.length} stale default models.`);
  return { refreshedModels, skippedModels };
}
