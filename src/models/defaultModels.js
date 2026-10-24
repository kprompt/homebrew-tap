/**
 * Module to manage default models with refresh capability.
 * This module exports functions to get and refresh default models,
 * ensuring models are not stale.
 */

const DEFAULT_MODEL_TTL_MS = 1000 * 60 * 60 * 24; // 24 hours TTL

// Simulated default models data source
let defaultModels = {
  modelA: { version: 1, data: { /* model data */ }, lastUpdated: Date.now() },
  modelB: { version: 1, data: { /* model data */ }, lastUpdated: Date.now() },
};

/**
 * Simulate fetching fresh model data from a remote source.
 * In real implementation, this would be an async call to a service or DB.
 * Here we simulate with a synchronous update for simplicity.
 */
function fetchFreshModelData(modelName) {
  // Simulate new version and data
  return {
    version: defaultModels[modelName].version + 1,
    data: { refreshed: true, timestamp: Date.now() },
    lastUpdated: Date.now(),
  };
}

/**
 * Returns the current default models.
 * If any model is stale, refreshes it before returning.
 */
function getDefaultModels() {
  const now = Date.now();
  Object.keys(defaultModels).forEach((modelName) => {
    const model = defaultModels[modelName];
    if (now - model.lastUpdated > DEFAULT_MODEL_TTL_MS) {
      // Refresh stale model
      defaultModels[modelName] = fetchFreshModelData(modelName);
    }
  });
  return defaultModels;
}

/**
 * Force refresh all default models regardless of staleness.
 * Returns the refreshed models.
 */
function refreshDefaultModels() {
  Object.keys(defaultModels).forEach((modelName) => {
    defaultModels[modelName] = fetchFreshModelData(modelName);
  });
  return defaultModels;
}

module.exports = {
  getDefaultModels,
  refreshDefaultModels,
};
