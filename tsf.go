package tsf

/*
#cgo LDFLAGS: -lm

#define TSF_STATIC 1
#define TSF_IMPLEMENTATION 1
#include "tsf.h"

*/
import "C"

type OutputMode int

const (
	ModeStereoInterleaved OutputMode = C.TSF_STEREO_INTERLEAVED
	ModeStereoUnweaved    OutputMode = C.TSF_STEREO_UNWEAVED
	ModeMono              OutputMode = C.TSF_MONO
)

type TSF struct {
	tsf *C.tsf
}

func New(font []byte) *TSF {
	return &TSF{tsf: C.tsf_load_memory(C.CBytes(font), C.int(len(font)))}
}

func NewFile(font string) *TSF {
	return &TSF{tsf: C.tsf_load_filename(C.CString(font))}
}

func (tsf *TSF) Close()       { C.tsf_close(tsf.tsf) }
func (tsf *TSF) Reset()       { C.tsf_reset(tsf.tsf) }
func (tsf *TSF) PresetCount() { C.tsf_get_presetcount(tsf.tsf) }
func (tsf *TSF) PresetIndex(bank int, preset int) {
	C.tsf_get_presetindex(tsf.tsf, C.int(bank), C.int(preset))
}
func (tsf *TSF) PresetName(preset int) string {
	return C.GoString(C.tsf_get_presetname(tsf.tsf, C.int(preset)))
}
func (tsf *TSF) BankPresetName(bank, preset int) string {
	return C.GoString(C.tsf_bank_get_presetname(tsf.tsf, C.int(bank), C.int(preset)))
}
func (tsf *TSF) SetOutput(mode OutputMode, sr int, gain float32) {
	C.tsf_set_output(tsf.tsf, C.enum_TSFOutputMode(mode), C.int(sr), C.float(gain))
}
func (tsf *TSF) NoteOn(preset, key int, vel float32) {
	C.tsf_note_on(tsf.tsf, C.int(preset), C.int(key), C.float(vel))
}
func (tsf *TSF) BankNoteOn(bank, preset, key int, vel float32) int {
	return int(C.tsf_bank_note_on(tsf.tsf, C.int(bank), C.int(preset), C.int(key), C.float(vel)))
}
func (tsf *TSF) NoteOff(preset, key int) {
	C.tsf_note_off(tsf.tsf, C.int(preset), C.int(key))
}
func (tsf *TSF) BankNoteOff(bank, preset, key int) int {
	return int(C.tsf_bank_note_off(tsf.tsf, C.int(bank), C.int(preset), C.int(key)))
}
func (tsf *TSF) NoteOffAll()           { C.tsf_note_off_all(tsf.tsf) }
func (tsf *TSF) ActiveVoiceCount() int { return int(C.tsf_active_voice_count(tsf.tsf)) }
func (tsf *TSF) RenderInt16(samples []int16) {
	C.tsf_render_short(tsf.tsf, (*C.short)(&samples[0]), C.int(len(samples)), 0)
}
func (tsf *TSF) RenderFloat32(samples []float32) {
	C.tsf_render_float(tsf.tsf, (*C.float)(&samples[0]), C.int(len(samples)), 0)
}
func (tsf *TSF) SetChannelPresetIndex(channel, preset int) {
	C.tsf_channel_set_presetindex(tsf.tsf, C.int(channel), C.int(preset))
}
func (tsf *TSF) SetChannelPresetNumber(channel, preset int) int {
	return int(C.tsf_channel_set_presetnumber(tsf.tsf, C.int(channel), C.int(preset), 0))
}
func (tsf *TSF) SetChannelBank(channel, bank int) {
	C.tsf_channel_set_bank(tsf.tsf, C.int(channel), C.int(bank))
}
func (tsf *TSF) SetChannelBankPreset(channel, bank, preset int) int {
	return int(C.tsf_channel_set_bank_preset(tsf.tsf, C.int(channel), C.int(bank), C.int(preset)))
}
func (tsf *TSF) SetChannelPan(channel int, pan float32) {
	C.tsf_channel_set_pan(tsf.tsf, C.int(channel), C.float(pan))
}
func (tsf *TSF) SetChannelVolume(channel int, volume float32) {
	C.tsf_channel_set_volume(tsf.tsf, C.int(channel), C.float(volume))
}
func (tsf *TSF) SetChannelPitchWheel(channel, pitch_wheel int) {
	C.tsf_channel_set_pitchwheel(tsf.tsf, C.int(channel), C.int(pitch_wheel))
}
func (tsf *TSF) SetChannelPitchRange(channel int, pitch_range int) {
	C.tsf_channel_set_pitchrange(tsf.tsf, C.int(channel), C.float(pitch_range))
}
func (tsf *TSF) SetChannelTuning(channel int, tuning float32) {
	C.tsf_channel_set_tuning(tsf.tsf, C.int(channel), C.float(tuning))
}
func (tsf *TSF) ChannelNoteOn(channel, key int, vel float32) {
	C.tsf_channel_note_on(tsf.tsf, C.int(channel), C.int(key), C.float(vel))
}
func (tsf *TSF) ChannelNoteOff(channel, key int) {
	C.tsf_channel_note_off(tsf.tsf, C.int(channel), C.int(key))
}
func (tsf *TSF) ChannelNoteOffAll(channel int) {
	C.tsf_channel_note_off_all(tsf.tsf, C.int(channel))
}
func (tsf *TSF) ChannelSoundsOffAll(channel int) {
	C.tsf_channel_sounds_off_all(tsf.tsf, C.int(channel))
}
func (tsf *TSF) ChannelMIDIControl(channel, controller, control_value int) {
	C.tsf_channel_midi_control(tsf.tsf, C.int(channel), C.int(controller), C.int(control_value))
}
func (tsf *TSF) ChannelPresetIndex(channel int) int {
	return int(C.tsf_channel_get_preset_index(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelPresetBank(channel int) int {
	return int(C.tsf_channel_get_preset_bank(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelPresetNumber(channel int) int {
	return int(C.tsf_channel_get_preset_number(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelPan(channel int) float32 {
	return float32(C.tsf_channel_get_pan(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelVolume(channel int) float32 {
	return float32(C.tsf_channel_get_volume(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelPitchWheel(channel int) float32 {
	return float32(C.tsf_channel_get_pitchwheel(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelPitchRange(channel int) float32 {
	return float32(C.tsf_channel_get_pitchrange(tsf.tsf, C.int(channel)))
}
func (tsf *TSF) ChannelTuning(channel int) float32 {
	return float32(C.tsf_channel_get_tuning(tsf.tsf, C.int(channel)))
}
