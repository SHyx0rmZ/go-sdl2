package sdl

// #include <SDL2/SDL_keycode.h>
import "C"

type KeyCode int32

const (
	KeyCodeToScanCodeMask = 1 << 30
)

const (
	KeyCodeUnknown KeyCode = 0

	KeyCodeReturn           KeyCode = '\r'
	KeyCodeEscape           KeyCode = '\033'
	KeyCodeBackspace        KeyCode = '\b'
	KeyCodeTab              KeyCode = '\t'
	KeyCodeSpace            KeyCode = ' '
	KeyCodeExclamationMark  KeyCode = '!'
	KeyCodeDoubleQuote      KeyCode = '"'
	KeyCodeHash             KeyCode = '#'
	KeyCodePercent          KeyCode = '%'
	KeyCodeDollar           KeyCode = '$'
	KeyCodeAmpersand        KeyCode = '&'
	KeyCodeQuote            KeyCode = '\''
	KeyCodeLeftParenthesis  KeyCode = '('
	KeyCodeRightParenthesis KeyCode = ')'
	KeyCodeAsterisk         KeyCode = '*'
	KeyCodePlus             KeyCode = '+'
	KeyCodeComma            KeyCode = ','
	KeyCodeMinus            KeyCode = '-'
	KeyCodePeriod           KeyCode = '.'
	KeyCodeSlash            KeyCode = '/'
	KeyCode0                KeyCode = '0'
	KeyCode1                KeyCode = '1'
	KeyCode2                KeyCode = '2'
	KeyCode3                KeyCode = '3'
	KeyCode4                KeyCode = '4'
	KeyCode5                KeyCode = '5'
	KeyCode6                KeyCode = '6'
	KeyCode7                KeyCode = '7'
	KeyCode8                KeyCode = '8'
	KeyCode9                KeyCode = '9'
	KeyCodeColon            KeyCode = ':'
	KeyCodeSemilon          KeyCode = ';'
	KeyCodeLess             KeyCode = '<'
	KeyCodeEquals           KeyCode = '='
	KeyCodeGreater          KeyCode = '>'
	KeyCodeQuestionMark     KeyCode = '?'
	KeyCodeAt               KeyCode = '@'

	KeyCodeLeftBracket  KeyCode = '['
	KeyCodeBackslash    KeyCode = '\\'
	KeyCodeRightBracket KeyCode = ']'
	KeyCodeCaret        KeyCode = '^'
	KeyCodeUnderscore   KeyCode = '_'
	KeyCodeBackquote    KeyCode = '`'
	KeyCodeA            KeyCode = 'a'
	KeyCodeB            KeyCode = 'b'
	KeyCodeC            KeyCode = 'c'
	KeyCodeD            KeyCode = 'd'
	KeyCodeE            KeyCode = 'e'
	KeyCodeF            KeyCode = 'f'
	KeyCodeG            KeyCode = 'g'
	KeyCodeH            KeyCode = 'h'
	KeyCodeI            KeyCode = 'i'
	KeyCodeJ            KeyCode = 'j'
	KeyCodeK            KeyCode = 'k'
	KeyCodeL            KeyCode = 'l'
	KeyCodeM            KeyCode = 'm'
	KeyCodeN            KeyCode = 'n'
	KeyCodeO            KeyCode = 'o'
	KeyCodeP            KeyCode = 'p'
	KeyCodeQ            KeyCode = 'q'
	KeyCodeR            KeyCode = 'r'
	KeyCodeS            KeyCode = 's'
	KeyCodeT            KeyCode = 't'
	KeyCodeU            KeyCode = 'u'
	KeyCodeV            KeyCode = 'v'
	KeyCodeW            KeyCode = 'w'
	KeyCodeX            KeyCode = 'x'
	KeyCodeY            KeyCode = 'y'
	KeyCodeZ            KeyCode = 'z'

	KeyCodeCapsLock = KeyCodeToScanCodeMask | KeyCode(ScanCodeCapsLock)

	KeyCodeF1  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF1)
	KeyCodeF2  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF2)
	KeyCodeF3  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF3)
	KeyCodeF4  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF4)
	KeyCodeF5  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF5)
	KeyCodeF6  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF6)
	KeyCodeF7  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF7)
	KeyCodeF8  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF8)
	KeyCodeF9  = KeyCodeToScanCodeMask | KeyCode(ScanCodeF9)
	KeyCodeF10 = KeyCodeToScanCodeMask | KeyCode(ScanCodeF10)
	KeyCodeF11 = KeyCodeToScanCodeMask | KeyCode(ScanCodeF11)
	KeyCodeF12 = KeyCodeToScanCodeMask | KeyCode(ScanCodeF12)

	KeyCodePrintScreen         = KeyCodeToScanCodeMask | KeyCode(ScanCodePrintScreen)
	KeyCodeScrollLock          = KeyCodeToScanCodeMask | KeyCode(ScanCodeScrollLock)
	KeyCodePause               = KeyCodeToScanCodeMask | KeyCode(ScanCodePause)
	KeyCodeInsert              = KeyCodeToScanCodeMask | KeyCode(ScanCodeInsert)
	KeyCodeHome                = KeyCodeToScanCodeMask | KeyCode(ScanCodeInsert)
	KeyCodePageUp              = KeyCodeToScanCodeMask | KeyCode(ScanCodePageUp)
	KeyCodeDelete      KeyCode = '\177'
	KeyCodeEnd                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeEnd)
	KeyCodePageDown            = KeyCodeToScanCodeMask | KeyCode(ScanCodePageDown)
	KeyCodeRight               = KeyCodeToScanCodeMask | KeyCode(ScanCodeRight)
	KeyCodeLeft                = KeyCodeToScanCodeMask | KeyCode(ScanCodeLeft)
	KeyCodeDown                = KeyCodeToScanCodeMask | KeyCode(ScanCodeDown)
	KeyCodeUp                  = KeyCodeToScanCodeMask | KeyCode(ScanCodeUp)

	KeyCodeNumLockClear   = KeyCodeToScanCodeMask | KeyCode(ScanCodeNumLockClear)
	KeyCodeKeyPadDivide   = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadDivide)
	KeyCodeKeyPadMultiply = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMultiply)
	KeyCodeKeyPadMinus    = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMinus)
	KeyCodeKeyPadPlus     = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadPlus)
	KeyCodeKeyPadEnter    = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadEnter)
	KeyCodeKeyPad1        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad1)
	KeyCodeKeyPad2        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad2)
	KeyCodeKeyPad3        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad3)
	KeyCodeKeyPad4        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad4)
	KeyCodeKeyPad5        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad5)
	KeyCodeKeyPad6        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad6)
	KeyCodeKeyPad7        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad7)
	KeyCodeKeyPad8        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad8)
	KeyCodeKeyPad9        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad9)
	KeyCodeKeyPad0        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad0)
	KeyCodeKeyPadPeriod   = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadPeriod)

	KeyCodeApplication       = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplication)
	KeyCodePower             = KeyCodeToScanCodeMask | KeyCode(ScanCodePower)
	KeyCodeKeyPadEquals      = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadEquals)
	KeyCodeF13               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF13)
	KeyCodeF14               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF14)
	KeyCodeF15               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF15)
	KeyCodeF16               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF16)
	KeyCodeF17               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF17)
	KeyCodeF18               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF18)
	KeyCodeF19               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF19)
	KeyCodeF20               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF20)
	KeyCodeF21               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF21)
	KeyCodeF22               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF22)
	KeyCodeF23               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF23)
	KeyCodeF24               = KeyCodeToScanCodeMask | KeyCode(ScanCodeF24)
	KeyCodeExecute           = KeyCodeToScanCodeMask | KeyCode(ScanCodeExecute)
	KeyCodeHelp              = KeyCodeToScanCodeMask | KeyCode(ScanCodeHelp)
	KeyCodeMenu              = KeyCodeToScanCodeMask | KeyCode(ScanCodeMenu)
	KeyCodeSelect            = KeyCodeToScanCodeMask | KeyCode(ScanCodeSelect)
	KeyCodeStop              = KeyCodeToScanCodeMask | KeyCode(ScanCodeStop)
	KeyCodeAgain             = KeyCodeToScanCodeMask | KeyCode(ScanCodeAgain)
	KeyCodeUndo              = KeyCodeToScanCodeMask | KeyCode(ScanCodeUndo)
	KeyCodeCut               = KeyCodeToScanCodeMask | KeyCode(ScanCodeCut)
	KeyCodeCopy              = KeyCodeToScanCodeMask | KeyCode(ScanCodeCopy)
	KeyCodePaste             = KeyCodeToScanCodeMask | KeyCode(ScanCodePaste)
	KeyCodeFind              = KeyCodeToScanCodeMask | KeyCode(ScanCodeFind)
	KeyCodeMute              = KeyCodeToScanCodeMask | KeyCode(ScanCodeMute)
	KeyCodeVolumeUp          = KeyCodeToScanCodeMask | KeyCode(ScanCodeVolumeUp)
	KeyCodeVolumeDown        = KeyCodeToScanCodeMask | KeyCode(ScanCodeVolumeDown)
	KeyCodeKeyPadComma       = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadComma)
	KeyCodeKeyPadEqualsAS400 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadEqualsAS400)

	KeyCodeAlternateErase = KeyCodeToScanCodeMask | KeyCode(ScanCodeAlternateErase)
	KeyCodeSystemRequest  = KeyCodeToScanCodeMask | KeyCode(ScanCodeSystemRequest)
	KeyCodeCancel         = KeyCodeToScanCodeMask | KeyCode(ScanCodeCancel)
	KeyCodeClear          = KeyCodeToScanCodeMask | KeyCode(ScanCodeClear)
	KeyCodePrior          = KeyCodeToScanCodeMask | KeyCode(ScanCodePrior)
	KeyCodeReturn2        = KeyCodeToScanCodeMask | KeyCode(ScanCodeReturn2)
	KeyCodeSeparator      = KeyCodeToScanCodeMask | KeyCode(ScanCodeSeparator)
	KeyCodeOut            = KeyCodeToScanCodeMask | KeyCode(ScanCodeOut)
	KeyCodeOper           = KeyCodeToScanCodeMask | KeyCode(ScanCodeOper)
	KeyCodeClearAgain     = KeyCodeToScanCodeMask | KeyCode(ScanCodeClearAgain)
	KeyCodeCrSel          = KeyCodeToScanCodeMask | KeyCode(ScanCodeCrSel)
	KeyCodeExSel          = KeyCodeToScanCodeMask | KeyCode(ScanCodeExSel)

	KeyCodeKeyPad00                = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad00)
	KeyCodeKeyPad000               = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPad000)
	KeyCodeThousandsSeparator      = KeyCodeToScanCodeMask | KeyCode(ScanCodeThousandsSeparator)
	KeyCodeDecimalSeparator        = KeyCodeToScanCodeMask | KeyCode(ScanCodeDecimalSeparator)
	KeyCodeCurrencyUnit            = KeyCodeToScanCodeMask | KeyCode(ScanCodeCurrencyUnit)
	KeyCodeCurrencySubUnit         = KeyCodeToScanCodeMask | KeyCode(ScanCodeCurrencySubUnit)
	KeyCodeKeyPadLeftParenthesis   = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadLeftParenthesis)
	KeyCodeKeyPadRightParenthesis  = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadRightParenthesis)
	KeyCodeKeyPadLeftBrace         = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadLeftBrace)
	KeyCodeKeyPadRightBrace        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadRightBrace)
	KeyCodeKeyPadTab               = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadTab)
	KeyCodeKeyPadBackspace         = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadBackspace)
	KeyCodeKeyPadA                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadA)
	KeyCodeKeyPadB                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadB)
	KeyCodeKeyPadC                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadC)
	KeyCodeKeyPadD                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadD)
	KeyCodeKeyPadE                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadE)
	KeyCodeKeyPadF                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadF)
	KeyCodeKeyPadXor               = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadXor)
	KeyCodeKeyPadPower             = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadPower)
	KeyCodeKeyPadPercent           = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadPercent)
	KeyCodeKeyPadLess              = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadLess)
	KeyCodeKeyPadGreater           = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadGreater)
	KeyCodeKeyPadAmpersand         = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadAmpersand)
	KeyCodeKeyPadDoubleAmpersand   = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadDoubleAmpersand)
	KeyCodeKeyPadVerticalBar       = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadVerticalBar)
	KeyCodeKeyPadDoubleVerticalBar = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadDoubleVerticalBar)
	KeyCodeKeyPadColon             = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadColon)
	KeyCodeKeyPadHash              = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadHash)
	KeyCodeKeyPadSpace             = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadSpace)
	KeyCodeKeyPadAt                = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadAt)
	KeyCodeKeyPadExclamationMark   = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadExclamationMark)
	KeyCodeKeyPadMemoryStore       = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemoryStore)
	KeyCodeKeyPadMemoryRecall      = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemoryRecall)
	KeyCodeKeyPadMemoryClear       = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemoryClear)
	KeyCodeKeyPadMemoryAdd         = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemoryAdd)
	KeyCodeKeyPadMemorySubtract    = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemorySubtract)
	KeyCodeKeyPadMemoryMultiply    = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemoryMultiply)
	KeyCodeKeyPadMemoryDivide      = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadMemoryDivide)
	KeyCodeKeyPadPlusMinus         = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadPlusMinus)
	KeyCodeKeyPadClear             = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadClear)
	KeyCodeKeyPadClearEntry        = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadClearEntry)
	KeyCodeKeyPadBinary            = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadBinary)
	KeyCodeKeyPadOctal             = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadOctal)
	KeyCodeKeyPadDecimal           = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadDecimal)
	KeyCodeKeyPadHexadecimal       = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyPadHexadecimal)

	KeyCodeLeftControl  = KeyCodeToScanCodeMask | KeyCode(ScanCodeLeftControl)
	KeyCodeLeftShift    = KeyCodeToScanCodeMask | KeyCode(ScanCodeLeftShift)
	KeyCodeLeftAlt      = KeyCodeToScanCodeMask | KeyCode(ScanCodeLeftAlt)
	KeyCodeLeftGUI      = KeyCodeToScanCodeMask | KeyCode(ScanCodeLeftGUI)
	KeyCodeRightControl = KeyCodeToScanCodeMask | KeyCode(ScanCodeRightControl)
	KeyCodeRightShift   = KeyCodeToScanCodeMask | KeyCode(ScanCodeRightShift)
	KeyCodeRightAlt     = KeyCodeToScanCodeMask | KeyCode(ScanCodeRightAlt)
	KeyCodeRightGUI     = KeyCodeToScanCodeMask | KeyCode(ScanCodeRightGUI)

	KeyCodeMode = KeyCodeToScanCodeMask | KeyCode(ScanCodeMode)

	KeyCodeAudioNext                   = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioNext)
	KeyCodeAudioPrevious               = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioPrevious)
	KeyCodeAudioStop                   = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioStop)
	KeyCodeAudioPlay                   = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioPlay)
	KeyCodeAudioMute                   = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioMute)
	KeyCodeMediaSelect                 = KeyCodeToScanCodeMask | KeyCode(ScanCodeMediaSelect)
	KeyCodeWWW                         = KeyCodeToScanCodeMask | KeyCode(ScanCodeWWW)
	KeyCodeMail                        = KeyCodeToScanCodeMask | KeyCode(ScanCodeMail)
	KeyCodeCalculator                  = KeyCodeToScanCodeMask | KeyCode(ScanCodeCalculator)
	KeyCodeComputer                    = KeyCodeToScanCodeMask | KeyCode(ScanCodeComputer)
	KeyCodeApplicationControlSearch    = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlSearch)
	KeyCodeApplicationControlHome      = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlHome)
	KeyCodeApplicationControlBack      = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlBack)
	KeyCodeApplicationControlForward   = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlForward)
	KeyCodeApplicationControlStop      = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlStop)
	KeyCodeApplicationControlRefresh   = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlRefresh)
	KeyCodeApplicationControlBookmarks = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplicationControlBookmarks)

	KeyCodeBrightnessDown             = KeyCodeToScanCodeMask | KeyCode(ScanCodeBrightnessDown)
	KeyCodeBrightnessUp               = KeyCodeToScanCodeMask | KeyCode(ScanCodeBrightnessUp)
	KeyCodeDisplaySwitch              = KeyCodeToScanCodeMask | KeyCode(ScanCodeDisplaySwitch)
	KeyCodeKeyboardIlluminationToggle = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyboardIlluminationToggle)
	KeyCodeKeyboardIlluminationDown   = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyboardIlluminationDown)
	KeyCodeKeyboardIlluminationUp     = KeyCodeToScanCodeMask | KeyCode(ScanCodeKeyboardIlluminationUp)
	KeyCodeEject                      = KeyCodeToScanCodeMask | KeyCode(ScanCodeEject)
	KeyCodeSleep                      = KeyCodeToScanCodeMask | KeyCode(ScanCodeSleep)
	KeyCodeApplication1               = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplication1)
	KeyCodeApplication2               = KeyCodeToScanCodeMask | KeyCode(ScanCodeApplication2)

	KeyCodeAudioRewind      = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioRewind)
	KeyCodeAudioFastForward = KeyCodeToScanCodeMask | KeyCode(ScanCodeAudioFastForward)
)

type KeyModifiers uint16

const (
	KeyModifierNone KeyModifiers = 0
)

const (
	KeyModifierLeftShift KeyModifiers = 0x0001 << iota
	KeyModifierRightShift
)

const (
	KeyModifierLeftControl KeyModifiers = 0x0040 << iota
	KeyModifierRightControl
	KeyModifierLeftAlt
	KeyModifierRightAlt
	KeyModifierLeftGUI
	KeyModifierRightGUI
	KeyModifierNumLock
	KeyModifierCapsLock
	KeyModifierMode
	KeyModifierReserved
)

const (
	KeyModifierControl = KeyModifierLeftControl | KeyModifierRightControl
	KeyModifierShift   = KeyModifierLeftShift | KeyModifierRightShift
	KeyModifierAlt     = KeyModifierLeftAlt | KeyModifierRightAlt
	KeyModifierGUI     = KeyModifierLeftGUI | KeyModifierRightGUI
)
