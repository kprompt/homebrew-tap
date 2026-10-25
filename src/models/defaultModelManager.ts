/**
 * Manages loading and refreshing of default models.
 */

import fs from 'fs';
import path from 'path';

const DEFAULT_MODEL_PATH = path.resolve(__dirname, '../../models/defaultModel.json');
const REFRESH_INTERVAL_MS = 1000 * 60 * 60 * 24; // 24 hours

interface DefaultModel {
  version: string;
  data: any;
  lastUpdated: number;
}

class DefaultModelManager {
  private model: DefaultModel | null = null;
  private lastLoadTime: number = 0;

  constructor() {
    this.loadModel();
    // Setup periodic refresh
    setInterval(() => {
      this.refreshIfStale();
    }, REFRESH_INTERVAL_MS);
  }

  private loadModel() {
    try {
      const raw = fs.readFileSync(DEFAULT_MODEL_PATH, 'utf-8');
      const parsed: DefaultModel = JSON.parse(raw);
      this.model = parsed;
      this.lastLoadTime = Date.now();
    } catch (err) {
      console.error('Failed to load default model:', err);
      this.model = null;
    }
  }

  private isStale(): boolean {
    if (!this.model) return true;
    // Consider stale if last loaded more than 24 hours ago
    return (Date.now() - this.lastLoadTime) > REFRESH_INTERVAL_MS;
  }

  refreshIfStale() {
    if (this.isStale()) {
      this.loadModel();
    }
  }

  getModel(): DefaultModel | null {
    this.refreshIfStale();
    return this.model;
  }
}

export const defaultModelManager = new DefaultModelManager();
