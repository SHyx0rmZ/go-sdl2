package sdl

// #include <SDL2/SDL_scancode.h>
import "C"

type ScanCode uint

const (
	ScanCodeUnknown ScanCode = 0
)

const (
	ScanCodeA ScanCode = 4 + iota
	ScanCodeB
	ScanCodeC
	ScanCodeD
	ScanCodeE
	ScanCodeF
	ScanCodeG
	ScanCodeH
	ScanCodeI
	ScanCodeJ
	ScanCodeK
	ScanCodeL
	ScanCodeM
	ScanCodeN
	ScanCodeO
	ScanCodeP
	ScanCodeQ
	ScanCodeR
	ScanCodeS
	ScanCodeT
	ScanCodeU
	ScanCodeV
	ScanCodeW
	ScanCodeX
	ScanCodeY
	ScanCodeZ

	ScanCode1
	ScanCode2
	ScanCode3
	ScanCode4
	ScanCode5
	ScanCode6
	ScanCode7
	ScanCode8
	ScanCode9
	ScanCode0

	ScanCodeReturn
	ScanCodeEscape
	ScanCodeBackspace
	ScanCodeTab
	ScanCodeSpace

	ScanCodeMinus
	ScanCodeEquals
	ScanCodeLeftBracket
	ScanCodeRightBracket
	ScanCodeBackslash
	ScanCodeNonUSHash
	ScanCodeSemicolon
	ScanCodeApostrophe
	ScanCodeGrave
	ScanCodeComma
	ScanCodePeriod
	ScanCodeSlash

	ScanCodeCapsLock

	ScanCodeF1
	ScanCodeF2
	ScanCodeF3
	ScanCodeF4
	ScanCodeF5
	ScanCodeF6
	ScanCodeF7
	ScanCodeF8
	ScanCodeF9
	ScanCodeF10
	ScanCodeF11
	ScanCodeF12

	ScanCodePrintScreen
	ScanCodeScrollLock
	ScanCodePause
	ScanCodeInsert
	ScanCodeHome
	ScanCodePageUp
	ScanCodeDelete
	ScanCodeEnd
	ScanCodePageDown
	ScanCodeRight
	ScanCodeLeft
	ScanCodeDown
	ScanCodeUp

	ScanCodeNumLockClear

	ScanCodeKeyPadDivide
	ScanCodeKeyPadMultiply
	ScanCodeKeyPadMinus
	ScanCodeKeyPadPlus
	ScanCodeKeyPadEnter
	ScanCodeKeyPad1
	ScanCodeKeyPad2
	ScanCodeKeyPad3
	ScanCodeKeyPad4
	ScanCodeKeyPad5
	ScanCodeKeyPad6
	ScanCodeKeyPad7
	ScanCodeKeyPad8
	ScanCodeKeyPad9
	ScanCodeKeyPad0
	ScanCodeKeyPadPeriod

	ScanCodeNonUSBackslash
	ScanCodeApplication
	ScanCodePower
	ScanCodeKeyPadEquals
	ScanCodeF13
	ScanCodeF14
	ScanCodeF15
	ScanCodeF16
	ScanCodeF17
	ScanCodeF18
	ScanCodeF19
	ScanCodeF20
	ScanCodeF21
	ScanCodeF22
	ScanCodeF23
	ScanCodeF24
	ScanCodeExecute
	ScanCodeHelp
	ScanCodeMenu
	ScanCodeSelect
	ScanCodeStop
	ScanCodeAgain
	ScanCodeUndo
	ScanCodeCut
	ScanCodeCopy
	ScanCodePaste
	ScanCodeFind
	ScanCodeMute
	ScanCodeVolumeUp
	ScanCodeVolumeDown
	_ // ScanCodeLockingCapsLock
	_ // ScanCodeLockingNumLock
	_ // ScanCodeLockingScrollLock
	ScanCodeKeyPadComma
	ScanCodeKeyPadEqualsAS400
	ScanCodeInternational1
	ScanCodeInternational2
	ScanCodeInternational3
	ScanCodeInternational4
	ScanCodeInternational5
	ScanCodeInternational6
	ScanCodeInternational7
	ScanCodeInternational8
	ScanCodeInternational9
	ScanCodeLanguage1
	ScanCodeLanguage2
	ScanCodeLanguage3
	ScanCodeLanguage4
	ScanCodeLanguage5
	ScanCodeLanguage6
	ScanCodeLanguage7
	ScanCodeLanguage8
	ScanCodeLanguage9

	ScanCodeAlternateErase
	ScanCodeSystemRequest
	ScanCodeCancel
	ScanCodeClear
	ScanCodePrior
	ScanCodeReturn2
	ScanCodeSeparator
	ScanCodeOut
	ScanCodeOper
	ScanCodeClearAgain
	ScanCodeCrSel
	ScanCodeExSel
)

const (
	ScanCodeKeyPad00 ScanCode = 176 + iota
	ScanCodeKeyPad000
	ScanCodeThousandsSeparator
	ScanCodeDecimalSeparator
	ScanCodeCurrencyUnit
	ScanCodeCurrencySubUnit
	ScanCodeKeyPadLeftParenthesis
	ScanCodeKeyPadRightParenthesis
	ScanCodeKeyPadLeftBrace
	ScanCodeKeyPadRightBrace
	ScanCodeKeyPadTab
	ScanCodeKeyPadBackspace
	ScanCodeKeyPadA
	ScanCodeKeyPadB
	ScanCodeKeyPadC
	ScanCodeKeyPadD
	ScanCodeKeyPadE
	ScanCodeKeyPadF
	ScanCodeKeyPadXor
	ScanCodeKeyPadPower
	ScanCodeKeyPadPercent
	ScanCodeKeyPadLess
	ScanCodeKeyPadGreater
	ScanCodeKeyPadAmpersand
	ScanCodeKeyPadDoubleAmpersand
	ScanCodeKeyPadVerticalBar
	ScanCodeKeyPadDoubleVerticalBar
	ScanCodeKeyPadColon
	ScanCodeKeyPadHash
	ScanCodeKeyPadSpace
	ScanCodeKeyPadAt
	ScanCodeKeyPadExclamationMark
	ScanCodeKeyPadMemoryStore
	ScanCodeKeyPadMemoryRecall
	ScanCodeKeyPadMemoryClear
	ScanCodeKeyPadMemoryAdd
	ScanCodeKeyPadMemorySubtract
	ScanCodeKeyPadMemoryMultiply
	ScanCodeKeyPadMemoryDivide
	ScanCodeKeyPadPlusMinus
	ScanCodeKeyPadClear
	ScanCodeKeyPadClearEntry
	ScanCodeKeyPadBinary
	ScanCodeKeyPadOctal
	ScanCodeKeyPadDecimal
	ScanCodeKeyPadHexadecimal
)

const (
	ScanCodeLeftControl ScanCode = 224 + iota
	ScanCodeLeftShift
	ScanCodeLeftAlt
	ScanCodeLeftGUI
	ScanCodeRightControl
	ScanCodeRightShift
	ScanCodeRightAlt
	ScanCodeRightGUI
)

const (
	ScanCodeMode ScanCode = 257 + iota

	ScanCodeAudioNext
	ScanCodeAudioPrevious
	ScanCodeAudioStop
	ScanCodeAudioPlay
	ScanCodeAudioMute
	ScanCodeMediaSelect
	ScanCodeWWW
	ScanCodeMail
	ScanCodeCalculator
	ScanCodeComputer
	ScanCodeApplicationControlSearch
	ScanCodeApplicationControlHome
	ScanCodeApplicationControlBack
	ScanCodeApplicationControlForward
	ScanCodeApplicationControlStop
	ScanCodeApplicationControlRefresh
	ScanCodeApplicationControlBookmarks

	ScanCodeBrightnessDown
	ScanCodeBrightnessUp
	ScanCodeDisplaySwitch
	ScanCodeKeyboardIlluminationToggle
	ScanCodeKeyboardIlluminationDown
	ScanCodeKeyboardIlluminationUp
	ScanCodeEject
	ScanCodeSleep

	ScanCodeApplication1
	ScanCodeApplication2

	ScanCodeAudioRewind
	ScanCodeAudioFastForward
)

const (
	NumScanCodes = 512
)
