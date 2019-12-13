package object

import (
	"testing"
)

func TestPhotosPhoto_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "photo20_10",
			fields: fields{10, 20},
			want:   "photo20_10",
		},
		{
			name:   "photo-10_20",
			fields: fields{20, -10},
			want:   "photo-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			photo := PhotosPhoto{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := photo.ToAttachment(); got != tt.want {
				t.Errorf("PhotosPhoto.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhotosPhotoFull_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "photo20_10",
			fields: fields{10, 20},
			want:   "photo20_10",
		},
		{
			name:   "photo-10_20",
			fields: fields{20, -10},
			want:   "photo-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			photo := PhotosPhotoFull{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := photo.ToAttachment(); got != tt.want {
				t.Errorf("PhotosPhotoFull.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhotosPhotoFullXtrRealOffset_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "photo20_10",
			fields: fields{10, 20},
			want:   "photo20_10",
		},
		{
			name:   "photo-10_20",
			fields: fields{20, -10},
			want:   "photo-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			photo := PhotosPhotoFullXtrRealOffset{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := photo.ToAttachment(); got != tt.want {
				t.Errorf("PhotosPhotoFullXtrRealOffset.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhotosPhotoXtrRealOffset_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "photo20_10",
			fields: fields{10, 20},
			want:   "photo20_10",
		},
		{
			name:   "photo-10_20",
			fields: fields{20, -10},
			want:   "photo-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			photo := PhotosPhotoXtrRealOffset{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := photo.ToAttachment(); got != tt.want {
				t.Errorf("PhotosPhotoXtrRealOffset.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_photosPhotoAlbum_ToAttachment(t *testing.T) {
	type fields struct {
		ID      string
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "album20_10",
			fields: fields{"10", 20},
			want:   "album20_10",
		},
		{
			name:   "album-10_20",
			fields: fields{"20", -10},
			want:   "album-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := photosPhotoAlbum{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := album.ToAttachment(); got != tt.want {
				t.Errorf("photosPhotoAlbum.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhotosPhotoAlbumFull_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "album20_10",
			fields: fields{10, 20},
			want:   "album20_10",
		},
		{
			name:   "album-10_20",
			fields: fields{20, -10},
			want:   "album-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			album := PhotosPhotoAlbumFull{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := album.ToAttachment(); got != tt.want {
				t.Errorf("PhotosPhotoAlbumFull.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhotosPhotoXtrTagInfo_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "photo20_10",
			fields: fields{10, 20},
			want:   "photo20_10",
		},
		{
			name:   "photo-10_20",
			fields: fields{20, -10},
			want:   "photo-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			photo := PhotosPhotoXtrTagInfo{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := photo.ToAttachment(); got != tt.want {
				t.Errorf("PhotosPhotoXtrTagInfo.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
