package bilibili

import (
	"testing"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/test_helper"
)

func TestCore_FetchMetaAndResourceInfo(t *testing.T) {
	type fields struct {
		Opts *args.Options
	}
	tests := []struct {
		name          string
		fields        fields
		wantMediaMeta *meta.MediaMeta
		wantErr       bool
	}{
		{
			fields: fields{Opts: &args.Options{Url: "https://www.bilibili.com/video/BV1eb4y187AG?spm_id_from=444.41.0.0"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "你永远叫不醒一只装睡的狗，除非去玩！",
				Description: "-",
				Artist:      "牧羊小队长",
				Album:       "Bilibili",
				Audio: meta.Audio{
					Url: ".m4s",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Core{
				Opts: tt.fields.Opts,
			}
			gotMediaMeta, err := c.FetchMetaAndResourceInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMetaAndResourceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			test_helper.TestMediaMeta(t, gotMediaMeta, tt.wantMediaMeta)
		})
	}
}