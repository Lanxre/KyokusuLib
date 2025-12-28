package static

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	UPLOAD_DIR string = "./uploads"
	UPLOAD_AVATAR_DIR string = "./uploads/avatars"
	UPLOAD_BANNER_DIR string = "./uploads/banners"
)

func CreateStaticDirs(r *mux.Router) error {
	if err := createStaticDir(); err != nil {
		return err
	}

	if err := createAvatarDir(); err != nil {
		return err
	}

	if err := createBannerDir(); err != nil {
		return err
	}

	r.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/", http.FileServer(http.Dir(UPLOAD_DIR))),
	)

	return nil
}


func createAvatarDir() error {
	if err := os.MkdirAll(UPLOAD_AVATAR_DIR, os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads avatars directory: %v", err)
	}

	return nil
}

func createStaticDir() error {
	if err := os.MkdirAll(UPLOAD_DIR, 0755); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}

	return nil
}

func createBannerDir() error {
	if err := os.MkdirAll(UPLOAD_BANNER_DIR, os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads avatars directory: %v", err)
	}

	return nil
}