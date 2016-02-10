package ytdl

import "strconv"

// FormatKey is a string type containing a key in a video format map
type FormatKey string

// Available format Keys
const (
	FormatExtensionKey     FormatKey = "ext"
	FormatResolutionKey    FormatKey = "res"
	FormatVideoEncodingKey FormatKey = "videnc"
	FormatAudioEncodingKey FormatKey = "audenc"
	FormatItagKey          FormatKey = "itag"
	FormatAudioBitrateKey  FormatKey = "audbr"
)

// Format is a youtube is a static youtube video format
type Format struct {
	Itag          int    `json:"itag"`
	Extension     string `json:"extension"`
	Resolution    string `json:"resolution"`
	VideoEncoding string `json:"videoEncoding"`
	AudioEncoding string `json:"audioEncoding"`
	AudioBitrate  int    `json:"audioBitrate"`
	meta          map[string]interface{}
}

func newFormat(itag int) (Format, bool) {
	if f, ok := FORMATS[itag]; ok {
		f.meta = make(map[string]interface{})
		return f, true
	}
	return Format{}, false
}

// ValueForKey gets the format value for a format key, used for filtering
func (f Format) ValueForKey(key FormatKey) interface{} {
	switch key {
	case FormatItagKey:
		return f.Itag
	case FormatExtensionKey:
		return f.Extension
	case FormatResolutionKey:
		return f.Resolution
	case FormatVideoEncodingKey:
		return f.VideoEncoding
	case FormatAudioEncodingKey:
		return f.AudioEncoding
	case FormatAudioBitrateKey:
		return f.AudioBitrate
	default:
		if f.meta != nil {
			return f.meta[string(key)]
		}
		return nil
	}
}

func (f Format) CompareKey(other Format, key FormatKey) int {
	switch key {
	case FormatResolutionKey:
		res := f.ValueForKey(key).(string)
		res1, res2 := 0, 0
		if res != "" {
			res1, _ = strconv.Atoi(res[0 : len(res)-2])
		}
		res = other.ValueForKey(key).(string)
		if res != "" {
			res2, _ = strconv.Atoi(res[0 : len(res)-2])
		}
		return res1 - res2
	case FormatAudioBitrateKey:
		return f.ValueForKey(key).(int) - other.ValueForKey(key).(int)
	default:
		return 0
	}
}

// FORMATS is a map of all itags and their formats
var FORMATS = map[int]Format{
	5: {
		Extension:     "flv",
		Resolution:    "240p",
		VideoEncoding: "Sorenson H.283",
		AudioEncoding: "mp3",
		Itag:          5,
		AudioBitrate:  64,
	},
	6: {
		Extension:     "flv",
		Resolution:    "270p",
		VideoEncoding: "Sorenson H.263",
		AudioEncoding: "mp3",
		Itag:          6,
		AudioBitrate:  64,
	},
	13: {
		Extension:     "3gp",
		Resolution:    "",
		VideoEncoding: "MPEG-4 Visual",
		AudioEncoding: "aac",
		Itag:          13,
		AudioBitrate:  0,
	},
	17: {
		Extension:     "3gp",
		Resolution:    "144p",
		VideoEncoding: "MPEG-4 Visual",
		AudioEncoding: "aac",
		Itag:          17,
		AudioBitrate:  24,
	},
	18: {
		Extension:     "mp4",
		Resolution:    "360p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          18,
		AudioBitrate:  96,
	},
	22: {
		Extension:     "mp4",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          22,
		AudioBitrate:  192,
	},
	34: {
		Extension:     "flv",
		Resolution:    "480p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          34,
		AudioBitrate:  128,
	},
	35: {
		Extension:     "flv",
		Resolution:    "360p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          35,
		AudioBitrate:  128,
	},
	36: {
		Extension:     "3gp",
		Resolution:    "240p",
		VideoEncoding: "MPEG-4 Visual",
		AudioEncoding: "aac",
		Itag:          36,
		AudioBitrate:  36,
	},
	37: {
		Extension:     "mp4",
		Resolution:    "1080p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          37,
		AudioBitrate:  192,
	},
	38: {
		Extension:     "mp4",
		Resolution:    "3072p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          38,
		AudioBitrate:  192,
	},
	43: {
		Extension:     "webm",
		Resolution:    "360p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          43,
		AudioBitrate:  128,
	},
	44: {
		Extension:     "webm",
		Resolution:    "480p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          44,
		AudioBitrate:  128,
	},
	45: {
		Extension:     "webm",
		Resolution:    "720p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          45,
		AudioBitrate:  192,
	},
	46: {
		Extension:     "webm",
		Resolution:    "1080p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          46,
		AudioBitrate:  192,
	},
	82: {
		Extension:     "mp4",
		Resolution:    "360p",
		VideoEncoding: "H.264",
		Itag:          82,
		AudioBitrate:  96,
	},
	83: {
		Extension:     "mp4",
		Resolution:    "240p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          83,
		AudioBitrate:  96,
	},
	84: {
		Extension:     "mp4",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          84,
		AudioBitrate:  192,
	},
	85: {
		Extension:     "mp4",
		Resolution:    "1080p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          85,
		AudioBitrate:  192,
	},
	100: {
		Extension:     "webm",
		Resolution:    "360p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          100,
		AudioBitrate:  128,
	},
	101: {
		Extension:     "webm",
		Resolution:    "360p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          101,
		AudioBitrate:  192,
	},
	102: {
		Extension:     "webm",
		Resolution:    "720p",
		VideoEncoding: "VP8",
		AudioEncoding: "vorbis",
		Itag:          102,
		AudioBitrate:  192,
	},
	// DASH (video only)
	133: {
		Extension:     "mp4",
		Resolution:    "240p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          133,
		AudioBitrate:  0,
	},
	134: {
		Extension:     "mp4",
		Resolution:    "360p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          134,
		AudioBitrate:  0,
	},
	135: {
		Extension:     "mp4",
		Resolution:    "480p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          135,
		AudioBitrate:  0,
	},
	136: {
		Extension:     "mp4",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          136,
		AudioBitrate:  0,
	},
	137: {
		Extension:     "mp4",
		Resolution:    "1080p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          137,
		AudioBitrate:  0,
	},
	138: {
		Extension:     "mp4",
		Resolution:    "2160p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          138,
		AudioBitrate:  0,
	},
	160: {
		Extension:     "mp4",
		Resolution:    "144p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          160,
		AudioBitrate:  0,
	},
	242: {
		Extension:     "webm",
		Resolution:    "240p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          242,
		AudioBitrate:  0,
	},
	243: {
		Extension:     "webm",
		Resolution:    "360p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          243,
		AudioBitrate:  0,
	},
	244: {
		Extension:     "webm",
		Resolution:    "480p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          244,
		AudioBitrate:  0,
	},
	247: {
		Extension:     "webm",
		Resolution:    "720p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          247,
		AudioBitrate:  0,
	},
	248: {
		Extension:     "webm",
		Resolution:    "1080p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          248,
		AudioBitrate:  9,
	},
	264: {
		Extension:     "mp4",
		Resolution:    "1440p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          264,
		AudioBitrate:  0,
	},
	266: {
		Extension:     "mp4",
		Resolution:    "2160p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          266,
		AudioBitrate:  0,
	},
	271: {
		Extension:     "webm",
		Resolution:    "1440p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          271,
		AudioBitrate:  0,
	},
	272: {
		Extension:     "webm",
		Resolution:    "2160p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          272,
		AudioBitrate:  0,
	},
	278: {
		Extension:     "webm",
		Resolution:    "144p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          278,
		AudioBitrate:  0,
	},
	298: {
		Extension:     "mp4",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          298,
		AudioBitrate:  0,
	},
	299: {
		Extension:     "mp4",
		Resolution:    "1080p",
		VideoEncoding: "H.264",
		AudioEncoding: "",
		Itag:          299,
		AudioBitrate:  0,
	},
	302: {
		Extension:     "webm",
		Resolution:    "720p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          302,
		AudioBitrate:  0,
	},
	303: {
		Extension:     "webm",
		Resolution:    "1080p",
		VideoEncoding: "VP9",
		AudioEncoding: "",
		Itag:          303,
		AudioBitrate:  0,
	},
	// DASH (audio only)
	139: {
		Extension:     "mp4",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "aac",
		Itag:          139,
		AudioBitrate:  48,
	},
	140: {
		Extension:     "mp4",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "aac",
		Itag:          140,
		AudioBitrate:  128,
	},
	141: {
		Extension:     "mp4",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "aac",
		Itag:          141,
		AudioBitrate:  256,
	},
	171: {
		Extension:     "webm",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "vorbis",
		Itag:          171,
		AudioBitrate:  128,
	},
	172: {
		Extension:     "webm",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "vorbis",
		Itag:          172,
		AudioBitrate:  192,
	},
	249: {
		Extension:     "webm",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "opus",
		Itag:          249,
		AudioBitrate:  50,
	},
	250: {
		Extension:     "webm",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "opus",
		Itag:          250,
		AudioBitrate:  70,
	},
	251: {
		Extension:     "webm",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "opus",
		Itag:          251,
		AudioBitrate:  160,
	},
	// Live streaming
	92: {
		Extension:     "ts",
		Resolution:    "240p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          92,
		AudioBitrate:  48,
	},
	93: {
		Extension:     "ts",
		Resolution:    "480p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          93,
		AudioBitrate:  128,
	},
	94: {
		Extension:     "ts",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          94,
		AudioBitrate:  128,
	},
	95: {
		Extension:     "ts",
		Resolution:    "1080p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          95,
		AudioBitrate:  256,
	},
	96: {
		Extension:     "ts",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          96,
		AudioBitrate:  256,
	},
	120: {
		Extension:     "flv",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          120,
		AudioBitrate:  128,
	},
	127: {
		Extension:     "ts",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "aac",
		Itag:          127,
		AudioBitrate:  96,
	},
	128: {
		Extension:     "ts",
		Resolution:    "",
		VideoEncoding: "",
		AudioEncoding: "aac",
		Itag:          128,
		AudioBitrate:  96,
	},
	132: {
		Extension:     "ts",
		Resolution:    "240p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          132,
		AudioBitrate:  48,
	},
	151: {
		Extension:     "ts",
		Resolution:    "720p",
		VideoEncoding: "H.264",
		AudioEncoding: "aac",
		Itag:          151,
		AudioBitrate:  24,
	},
}
