package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TileProvider 接口现在需要一个 theme 参数
type TileProvider interface {
	GetTile(theme string, z, x, y int) ([]byte, error)
}

func newTileProvider(config []SourceConfig) map[string]TileProvider {
	var err error
	providers := make(map[string]TileProvider)
	for _, sourceConfig := range config {
		var provider TileProvider
		switch sourceConfig.Type {
		case "mbtiles":
			provider, err = NewMBTilesProvider(sourceConfig.Path)
		case "dir":
			provider, err = NewDirProvider(sourceConfig.Path)
		default:
			log.Printf("Warning: Unknown source type '%s' for source '%s'. Skipping.", sourceConfig.Type, sourceConfig.Name)
			continue
		}

		if err != nil {
			log.Printf("Error initializing source '%s': %v. Skipping.", sourceConfig.Name, err)
			continue
		}
		providers[sourceConfig.Name] = provider
		if _, ok := providers["default"]; !ok {
			providers["default"] = provider
		}
		log.Printf("Successfully loaded source '%s' (type: %s)", sourceConfig.Name, sourceConfig.Type)
	}

	if len(providers) == 0 {
		log.Fatalf("No valid tile sources were loaded. Exiting.")
	}
	return providers
}

// DirProvider 实现 TileProvider 接口（用于单主题目录）
// 同样忽略 theme 参数
type DirProvider struct {
	rootDir string
}

func NewDirProvider(path string) (*DirProvider, error) {
	// ... (此函数保持不变)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("tile directory does not exist: %s", path)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("path is not a directory: %s", path)
	}
	return &DirProvider{rootDir: path}, nil
}

func (p *DirProvider) GetTile(_ string, z, x, y int) ([]byte, error) {
	// 使用 _ 忽略 theme 参数
	// ... (此函数逻辑保持不变)
	tilePath := filepath.Join(p.rootDir, fmt.Sprintf("%d", z), fmt.Sprintf("%d", x), fmt.Sprintf("%d.png", y))
	data, err := os.ReadFile(tilePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MBTilesProvider 实现 TileProvider 接口
// 它没有主题概念，所以会忽略 theme 参数
type MBTilesProvider struct {
	db *sql.DB
}

func NewMBTilesProvider(path string) (*MBTilesProvider, error) {
	// ... (此函数保持不变)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open mbtiles file %s: %w", path, err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping mbtiles database %s: %w", path, err)
	}
	return &MBTilesProvider{db: db}, nil
}

func (p *MBTilesProvider) GetTile(_ string, z, x, y int) ([]byte, error) {
	// 使用 _ 忽略 theme 参数
	// ... (此函数逻辑保持不变)
	tmsY := (1 << uint(z)) - 1 - y
	var tileData []byte
	query := `SELECT tile_data FROM tiles WHERE zoom_level = ? AND tile_column = ? AND tile_row = ?`
	err := p.db.QueryRow(query, z, x, tmsY).Scan(&tileData)
	if err != nil {
		return nil, err
	}
	return tileData, nil
}
