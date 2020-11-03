package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

// NewAssetManager creates and assigns all necessary dependencies for the AssetManager top-level functions to work correctly
func NewAssetManager() (*AssetManager, error) {
	loader, err := d2loader.NewLoader(d2util.LogLevelDefault)
	if err != nil {
		return nil, err
	}

	records, err := d2records.NewRecordManager(d2util.LogLevelDebug)
	if err != nil {
		return nil, err
	}

	manager := &AssetManager{
		logger:     d2util.NewLogger(),
		loader:     loader,
		tables:     d2cache.CreateCache(tableBudget),
		animations: d2cache.CreateCache(animationBudget),
		fonts:      d2cache.CreateCache(fontBudget),
		palettes:   d2cache.CreateCache(paletteBudget),
		transforms: d2cache.CreateCache(paletteTransformBudget),
		Records:    records,
	}

	manager.logger.SetPrefix(logPrefix)

	err = manager.init()

	return manager, err
}
