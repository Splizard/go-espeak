/*
Go Bindings for espeak.
*/
package espeak

/*
#include <espeak/speak_lib.h>
#include <stdlib.h>
#include <string.h>

void* user_data;
unsigned int *unique_identifier;

static int EspeakInit()
{
		return (int) espeak_Initialize(0, 500, NULL, 0); 
}

static int EspeakSynth(const char *text, unsigned int position, espeak_POSITION_TYPE position_type, unsigned int end_position)
{
		unsigned int size_t;
		size_t = strlen(text)+1;
		return (int) espeak_Synth( text, size_t, position, position_type, end_position, 1,
    unique_identifier, user_data );
}

*/
// #cgo LDFLAGS: -lespeak
import "C"
import "unsafe"

type EspeakAudioOutput int

const (
	AUDIO_OUTPUT_PLAYBACK       EspeakAudioOutput = 0
	AUDIO_OUTPUT_RETRIEVAL      EspeakAudioOutput = 1
	AUDIO_OUTPUT_SYNCHRONOUS    EspeakAudioOutput = 2
	AUDIO_OUTPUT_SYNCH_PLAYBACK EspeakAudioOutput = 3
)

type EspeakPositionType int

const (
	POS_CHARACTER EspeakPositionType = 1
	POS_WORD      EspeakPositionType = 2
	POS_SENTENCE  EspeakPositionType = 3
)

func Init() int {
	return int(C.EspeakInit())
}

func Initialize(output EspeakAudioOutput, buflength int, path string, options int) int {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return int(C.espeak_Initialize(C.espeak_AUDIO_OUTPUT(output), C.int(buflength), cpath, C.int(options)))
}

func SetVoiceByName(voice string) int {
	cvoice := C.CString(voice)
	defer C.free(unsafe.Pointer(cvoice))
	return int(C.espeak_SetVoiceByName(cvoice))
}

func Synth(text string, position uint, positionType EspeakPositionType, endPosition uint) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return int(C.EspeakSynth(ctext, C.uint(position), C.espeak_POSITION_TYPE(positionType), C.uint(endPosition)))
}

func Say(text string) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return int(C.EspeakSynth(ctext, C.uint(0), C.espeak_POSITION_TYPE(0), C.uint(0)))
}

func Synchronize() int {
	return int(C.espeak_Synchronize())
}

func Sync() int {
	return int(C.espeak_Synchronize())
}

type EspeakParameter int

const (
	SILENCE     EspeakParameter = 0
	RATE        EspeakParameter = 1
	VOLUME      EspeakParameter = 2
	PITCH       EspeakParameter = 3
	RANGE       EspeakParameter = 4
	PUNCTUATION EspeakParameter = 5
	CAPITALS    EspeakParameter = 6
	WORDGAP     EspeakParameter = 7
	OPTIONS     EspeakParameter = 8
	INTONATION  EspeakParameter = 9

	RESERVED1      EspeakParameter = 10
	RESERVED2      EspeakParameter = 11
	EMPHASIS       EspeakParameter = 12
	LINELENGTH     EspeakParameter = 13
	VOICETYPE      EspeakParameter = 14
	N_SPEECH_PARAM EspeakParameter = 15
)

func SetParameter(parameter EspeakParameter, value int, relative int) int {
	return int(C.espeak_SetParameter(C.espeak_PARAMETER(parameter), C.int(value), C.int(relative)))
}

func GetParameter(parameter EspeakParameter) int {
	return int(C.espeak_GetParameter(C.espeak_PARAMETER(parameter),1))
}

func GetDefaultParameter(parameter EspeakParameter) int {
	return int(C.espeak_GetParameter(C.espeak_PARAMETER(parameter),0))
}

func Cancel() int {
	return int(C.espeak_Cancel())
}

func Terminate() int {
	return int(C.espeak_Terminate())
}

func IsPlaying() int {
	return int(C.espeak_IsPlaying())
}
