/**
 * Module to refresh stale default models.
 * 
 * This module exports a function `refreshStaleDefaultModels` which checks all default models
 * and refreshes those that are stale based on a configurable staleness threshold.
 */

import { getDefaultModels, updateModel } from './modelStore';
import { logger } from './logger';

export interface Model {
  id: string;
  name: string;
  lastRefreshed: number; // timestamp in ms
  isDefault: boolean;
  data: any;
}

export interface RefreshOptions {
  staleThresholdMs: number; // models older than this are considered stale
}

/**
 * Refreshes stale default models.
 * 
 * @param options - options including staleThresholdMs
 * @returns array of refreshed model ids
 */
export async function refreshStaleDefaultModels(
  options: RefreshOptions = { staleThresholdMs: 1000 * 60 * 60 * 24 } // default 24h
): Promise<string[]> {
  const now = Date.now();
  const models = await getDefaultModels();

  const staleModels = models.filter(
    (m) => m.isDefault && now - m.lastRefreshed > options.staleThresholdMs
  );

  const refreshedModelIds: string[] = [];

  for (const model of staleModels) {
    try {
      // Simulate refresh logic - e.g. fetch new data or re-generate model
      const refreshedData = await refreshModelData(model);
      const updatedModel = {
        ...model,
        data: refreshedData,
        lastRefreshed: now,
      };
      await updateModel(updatedModel);
      refreshedModelIds.push(model.id);
      logger.info(`Refreshed default model ${model.name} (${model.id})`);
    } catch (err) {
      logger.error(`Failed to refresh model ${model.id}: ${err}`);
    }
  }

  return refreshedModelIds;
}

/**
 * Simulate refreshing model data.
 * In real implementation, this would fetch or compute fresh model data.
 */
async function refreshModelData(model: Model): Promise<any> {
  // For demonstration, just return the existing data with a refreshed timestamp
  return {
    ...model.data,
    refreshedAt: Date.now(),
  };
}
