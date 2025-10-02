package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Serve(port int, handler *TileHandler) {
	r := chi.NewRouter()
	// 带主题的瓦片路由: /{theme}/{z}/{x}/{y}.png
	r.Get("/{theme}/{z}/{x}/{y}.png", handler.handleThemedTile)

	// --- 3. 启动服务器 ---
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

type TileHandler struct {
	provider map[string]TileProvider
}

func NewTileHandler(provider map[string]TileProvider) *TileHandler {
	return &TileHandler{provider: provider}
}

// handleThemedTile 处理带主题的瓦片请求
func (h *TileHandler) handleThemedTile(w http.ResponseWriter, r *http.Request) {
	theme := chi.URLParam(r, "theme")
	z, _ := strconv.Atoi(chi.URLParam(r, "z"))
	x, _ := strconv.Atoi(chi.URLParam(r, "x"))
	y, _ := strconv.Atoi(chi.URLParam(r, "y"))

	provider, ok := h.provider[theme]
	if !ok {
		provider = h.provider["default"]
	}

	tileData, err := provider.GetTile(theme, z, x, y)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "no such file") ||
			strings.Contains(err.Error(), "not found") {
			http.NotFound(w, r)
		} else {
			// chi 的 Logger 中间件已经记录了请求，这里只需记录错误
			log.Printf("Error getting tile: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	if _, err := w.Write(tileData); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
