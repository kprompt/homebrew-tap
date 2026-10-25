import { expect } from 'chai';
import fs from 'fs/promises';
import path from 'path';
import sinon from 'sinon';

import {
  loadDefaultModel,
  saveDefaultModel,
  isModelStale,
  fetchFreshDefaultModel,
  refreshDefaultModelIfStale,
  DefaultModel,
} from '../../src/models/defaultModelRefresher';

const MODEL_DIR = path.resolve(__dirname, '../../models');
const DEFAULT_MODEL_FILE = path.join(MODEL_DIR, 'defaultModel.json');

describe('defaultModelRefresher', () => {
  afterEach(async () => {
    // Cleanup model file after each test
    try {
      await fs.unlink(DEFAULT_MODEL_FILE);
    } catch {
      // ignore if file does not exist
    }
  });

  describe('loadDefaultModel', () => {
    it('returns null if file does not exist', async () => {
      const model = await loadDefaultModel();
      expect(model).to.be.null;
    });

    it('loads a valid model from disk', async () => {
      const testModel: DefaultModel = {
        version: 'v1',
        data: { foo: 'bar' },
        lastUpdated: Date.now(),
      };
      await saveDefaultModel(testModel);
      const loaded = await loadDefaultModel();
      expect(loaded).to.deep.equal(testModel);
    });
  });

  describe('isModelStale', () => {
    it('returns true if model is older than threshold', () => {
      const oldModel: DefaultModel = {
        version: 'v1',
        data: {},
        lastUpdated: Date.now() - 1000 * 60 * 60 * 24 * 8, // 8 days ago
      };
      expect(isModelStale(oldModel)).to.be.true;
    });

    it('returns false if model is recent', () => {
      const recentModel: DefaultModel = {
        version: 'v1',
        data: {},
        lastUpdated: Date.now() - 1000 * 60 * 60 * 24 * 2, // 2 days ago
      };
      expect(isModelStale(recentModel)).to.be.false;
    });
  });

  describe('refreshDefaultModelIfStale', () => {
    it('fetches fresh model if none exists', async () => {
      const freshModel = await refreshDefaultModelIfStale();
      expect(freshModel).to.have.property('version');
      expect(freshModel).to.have.property('data');
      expect(freshModel).to.have.property('lastUpdated');
      // Confirm file saved
      const savedModel = await loadDefaultModel();
      expect(savedModel).to.deep.equal(freshModel);
    });

    it('fetches fresh model if existing model is stale', async () => {
      const staleModel: DefaultModel = {
        version: 'v1',
        data: { stale: true },
        lastUpdated: Date.now() - 1000 * 60 * 60 * 24 * 8, // 8 days ago
      };
      await saveDefaultModel(staleModel);
      const freshModel = await refreshDefaultModelIfStale();
      expect(freshModel.version).to.not.equal(staleModel.version);
      expect(freshModel.lastUpdated).to.be.greaterThan(staleModel.lastUpdated);
    });

    it('returns existing model if not stale', async () => {
      const recentModel: DefaultModel = {
        version: 'v1',
        data: { fresh: true },
        lastUpdated: Date.now(),
      };
      await saveDefaultModel(recentModel);
      const model = await refreshDefaultModelIfStale();
      expect(model).to.deep.equal(recentModel);
    });
  });
});
