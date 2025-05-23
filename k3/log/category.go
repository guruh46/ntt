package log

import "strings"

// Category descripes all event classes available in K3 runtime.
type Category string

// ID returns the 4 letter code identifing K3 runtime log event
func (c Category) ID() string {
	return string(c)[0:4]
}

// IsError returns true if c is describes an error event.
func (c Category) IsError() bool {
	return strings.ToUpper(c.ID()) == c.ID()
}

// Len returns the total number of fields an event has.
func (c Category) Len() int {
	return 2 + strings.Count(string(c), "|")
}

// String returns an one-line description of the event class
func (c Category) String() string {
	return string(c)
}

const (
	ACFG    Category = "acfg|Assembly Configuration Info|Key|Value"
	ALEN    Category = "alen|Alt enter"
	ALLV    Category = "allv|Alt leave"
	ALRP    Category = "alrp|Alt repeat"
	ALWT    Category = "alwt|Alternation wait: no alternative ready|Time at which to wake unless an alternative becomes ready"
	ASEN    Category = "asen|Altstep enter|Altstep name + parameters"
	ASLV    Category = "aslv|Altstep leave|Altstep name + parameters"
	BCTR    Category = "bctr|Backtrace event tracks frames visited by components|Component name and tracked frames"
	COAL    Category = "coal|alive operation|Component name|Result: alive/killed"
	COCR    Category = "cocr|Component created|Component name|Alive-type indicator: alive/once"
	CODO    Category = "codo|Evaluate component.done()|Component name|Outcome"
	COFI    Category = "cofi|Component finished executing behaviour|Final verdict of component"
	COKD    Category = "cokd|Evaluate component.killed()|Component name|Outcome"
	COKI    Category = "coki|Component killed|Component name"
	CORU    Category = "coru|running operation on component|Component name|Result: running/stopped"
	COSP    Category = "cosp|Component stopped|Component name"
	COST    Category = "cost|Component started|Name of behaviour function"
	CPEN    Category = "cpen|Control part enter|Control part name"
	CPLV    Category = "cplv|Control part leave|Control part name"
	DBG1    Category = "dbg1|Debug1|message"
	DBG2    Category = "dbg2|Debug2|message"
	DBG3    Category = "dbg3|Debug3|message"
	DBUG    Category = "dbug|only for debug purpose|debug string"
	DECO    Category = "deco|Decoded a message|Type of message|Encode attribute|Variant attribute|Extension attribute"
	DECV    Category = "decv|decvalue|decoded value"
	DTAC    Category = "dtac|Default activate|Altstep invocation"
	DTDE    Category = "dtde|Default deactivate|Altstep invocation"
	DTEN    Category = "dten|Default enter|Default name + parameters"
	DTLV    Category = "dtlv|Default leave|Default name + parameters"
	DUMP    Category = "dump|Catalog item dump|item name|item detail"
	ELSE    Category = "else|Evaluation of else clause|Always ready"
	ENCO    Category = "enco|Encoded a message|Type of message|Encode attribute|Variant attribute|Extension attribute"
	FNEN    Category = "fnen|Function enter|Function name + parameters"
	FNLV    Category = "fnlv|Function leave|Function name + parameters + optional return value"
	FXEN    Category = "fxen|External function enter|External function name + parameters"
	FXLV    Category = "fxlv|External function leave|External function name + parameters + optional return value"
	GETV    Category = "getv|getverdict operation|Current verdict"
	ILEN    Category = "ilen|Interleave enter"
	ILLV    Category = "illv|Interleave leave"
	MATC    Category = "matc|Left-hand-side:value|Right-hand-side:template|List of mismatches (empty means matched)"
	MPAR    Category = "mpar|module parameter|name|value"
	PAON    Category = "paon|produces no log-line. To be used for switching on parameter logging for functions, altsteps, testcases"
	PERF    Category = "perf|Performance Measurement|key|value"
	PLLG    Category = "pllg|plugin related log message"
	PLOD    Category = "plod|Plugin loaded|Path to plugin|Plugin name|Plugin type"
	PTCK    Category = "ptck|Evaluate port.check()|Port name|Match template|Outcome"
	PTCL    Category = "ptcl|Port cleared|Port name"
	PTCN    Category = "ptcn|Port connected|First port name|Second port name"
	PTDI    Category = "ptdi|Port disconnected|First port name or all|Optional second port name"
	PTDS    Category = "ptds|Item discarded at port|Port name|Item detail (message, call, reply, exception) + value|Reason for discard"
	PTHA    Category = "ptha|Port halted|Port name"
	PTMP    Category = "ptmp|Port mapped|Component port name|System port name"
	PTPU    Category = "ptpu|Port published to external connector|Port name"
	PTQU    Category = "ptqu|Item queued to port|Port name|Item detail (message, call, reply, exception) + value"
	PTRE    Category = "ptre|Evaluate port.receive() on empty queue|[parameter name->]Port name|Match template"
	PTRX    Category = "ptrx|Evaluate port.receive()|[parameter name->]Port name|Match template|Outcome"
	PTSD    Category = "ptsd|Port send|Component port|System port|Message type name|Message value"
	PTSP    Category = "ptsp|Port stopped|Port name"
	PTST    Category = "ptst|Port started|Port name"
	PTTR    Category = "pttr|Evaluate port.trigger()|Port name|Match template|Outcome"
	PTUN    Category = "ptun|Port unmapped|Component port name or all|Optional system port name"
	RVON    Category = "rvon|produces value return. To be used for switching on return value logging for functions"
	SDBG    Category = "sdbg|SnapDebug|message"
	SETV    Category = "setv|setverdict operation|Previous verdict|New verdict|Reason"
	TCEN    Category = "tcen|Testcase enter|Testcase name + parameters"
	TCFI    Category = "tcfi|Testcase finished|testcase name|verdict"
	TCLV    Category = "tclv|Testcase leave|Testcase name + parameters"
	TCST    Category = "tcst|Testcase start|testcase name|guard duration"
	TMRD    Category = "tmrd|read operation|Timer name|Expired duration"
	TMRU    Category = "tmru|running operation for timer|Timer name|Result: false/true"
	TMSP    Category = "tmsp|Timer stop|Timer name"
	TMST    Category = "tmst|Timer start|Timer name|Timer duration"
	TMTO    Category = "tmto|Evaluate timer.timeout()|Timer name|Outcome"
	UACT    Category = "uact|User action|User message"
	ULOG    Category = "ulog|User action|User log"
	VACH    Category = "vach|Value changed|Type name|Value name|New value"
	VRSN    Category = "vrsn|Program version information|version number as X.Y.Z"
	WAIT    Category = "wait|Wait (real-time)|Time at which to wake"
	ErrACDC Category = "ACDC|Ambiguous Codec found|Name of conflicted type|Name of conflicting type"
	ErrARGI Category = "ARGI|Insufficient arguments"
	ErrARGS Category = "ARGS|Spurious (unexpected) argument|Text of the unexpected argument"
	ErrARRV Category = "ARRV|Array bounds violation|Value name|Erroneous array index"
	ErrARTE Category = "ARTE|More than one route possible for transmit operation (ambiguous)"
	ErrARUN Category = "ARUN|Already running|Identity of already running thread"
	ErrBREF Category = "BREF|Broken reference: indicates a reportable runtime bug|Erroneous offset"
	ErrCONE Category = "CONE|Connection error|First port name|Second port name|Error detail"
	ErrCONV Category = "CONV|Constraint violation|Type name|Constraint|Value name|Erroneous value"
	ErrDEAD Category = "DEAD|Deadlock: no alternative available"
	ErrDIRE Category = "DIRE|Discard report at port|Port name|Error detail"
	ErrDIV0 Category = "DIV0|Divide by zero|Value name"
	ErrDKEY Category = "DKEY|Duplicated map key value inside assignment list|Partial assignment list|Key Value"
	ErrDOME Category = "DOME|Domain error|Value name"
	ErrDTDE Category = "DTDE|Default deactivate error|Default state"
	ErrFILE Category = "FILE|Failed to open a file|Named of file"
	ErrFLOW Category = "FLOW|Control flow error: indicates a reportable bug|Erroneous operation"
	ErrFOEX Category = "FOEX|Report a foreign exception during operation: indicates a bug to be reported|C++ RTTI information when available|Error detail"
	ErrICST Category = "ICST|Content of the template|Error detail"
	ErrIGRP Category = "IGRP|Invalid group number in pattern match|Description of pattern|Requested invalid group number"
	ErrINOP Category = "INOP|Invalid operation|Operation detail"
	ErrLENE Category = "LENE|Length mismatch|Left value name|Left value length|Right value name|Right value length"
	ErrLIDE Category = "LIDE|Literal data error|Data type|Data in error"
	ErrLKUP Category = "LKUP|Lookup error (plugin system): no plugin offered requested service|Type of lookup that failed|First (or only) part of requested service|Second (optional) part of requested service"
	ErrMALF Category = "MALF|Malformed argument value|Erroneous option or argument|Erroneous value"
	ErrMANY Category = "MANY|Many valued|Value name"
	ErrMAPE Category = "MAPE|Mapping error|Component port name|System port name|Error detail"
	ErrNAME Category = "NAME|Broken name: indicates a reportable runtime bug|Erroneous name"
	ErrNOIM Category = "NOIM|Feature not implemented|Feature detail"
	ErrNRTE Category = "NRTE|No route found for transmit operation|Missing destination"
	ErrNSPR Category = "NSPR|No system port reference in unmap statement with parameters"
	ErrNULL Category = "NULL|Null value|Value name"
	ErrOMIT Category = "OMIT|Omitted value|Value name"
	ErrOPTM Category = "OPTM|Missing argument to option|Short or long form of the option name"
	ErrOPTS Category = "OPTS|Spurious (unexpected) argument to option|Short or long form of the option name"
	ErrOPTU Category = "OPTU|Unknown option|Short or long form of the option name"
	ErrOSUF Category = "OSUF|Object stack underflow (bug)"
	ErrPABV Category = "PABV|Parameter bound violation, probably from Codec|Erroneous index"
	ErrPAOV Category = "PAOV|Parameter 'out' violation: out or inout parameter not allowed|Name of erroneous parameter"
	ErrPARE Category = "PARE|Parse error (read value from string)|Value name|Erroneous input"
	ErrPATE Category = "PATE|Pattern error during parsing|Erroneous pattern|Error detail"
	ErrPLEX Category = "PLEX|Report a exception during plugin invocation: indicates a bug to be reported|Error context|Error detail"
	ErrPLOD Category = "PLOD|Plugin load error|Path to plugin|Plugin name"
	ErrPTFA Category = "PTFA|Port failed in external connector|Failure detail"
	ErrRNGE Category = "RNGE|Range error|Erroneous value"
	ErrSIZE Category = "SIZE|T3XF file size error: size not exact multiple of instruction width (bug)|Name of erroneous file"
	ErrSNAC Category = "SNAC|Starting behaviour on non-alive (killed) component|Component name"
	ErrSNAP Category = "SNAP|Invalid operation during snapshot"
	ErrSRUN Category = "SRUN|Start already running|Identity of already running thread"
	ErrSYSE Category = "SYSE|Report a std::system_error exception during operation: indicates a bug to be reported|error code|error description"
	ErrTIME Category = "TIME|Testcase execution time limit exceeded"
	ErrTSTP Category = "TSTP|Test stopped by the testcase.stop operation|Detail"
	ErrTYPE Category = "TYPE|Type mismatch (bug)|Expected type|Given type"
	ErrUBLK Category = "UBLK|Unterminated block in input file (corrupt file or compiler bug)|Offset of start of block|Offset of scan limit"
	ErrUNAS Category = "UNAS|Unassignable|Value name"
	ErrUNDF Category = "UNDF|Undefined value|Value name"
	ErrUNOP Category = "UNOP|Unkown operation: invalid operation encountered (corrupt file?)"
	ErrUTF8 Category = "UTF8|UTF-8 parse error|Error detail"
	ErrUTOB Category = "UTOB|Attempt to call a bound (runs on) behaviour from an unbound behaviour"
	ErrVRSN Category = "VRSN|Input file version error|Erroneous version"
	ErrWAIT Category = "WAIT|Time for a wait operation has already passed"
	ErrWCPA Category = "WCPA|The parameter passed to checkstate is not valid|Passed parameter value"
	ErrWIDE Category = "WIDE|Instruction width too wide (> 64-bit)|Name of the rejected file"
	ErrWPAC Category = "WPAC|Wrong number of parameters passed to paramaterized map or unmap statement|Expected number of parameters|Received number of parameters"
)

var (
	// Categories is a map of all categories. The key is the category name
	// and the description of the category.
	Categories = map[string]Category{
		"acfg": ACFG,
		"alen": ALEN,
		"allv": ALLV,
		"alrp": ALRP,
		"alwt": ALWT,
		"asen": ASEN,
		"aslv": ASLV,
		"bctr": BCTR,
		"coal": COAL,
		"cocr": COCR,
		"codo": CODO,
		"cofi": COFI,
		"cokd": COKD,
		"coki": COKI,
		"coru": CORU,
		"cosp": COSP,
		"cost": COST,
		"cpen": CPEN,
		"cplv": CPLV,
		"dbg1": DBG1,
		"dbg2": DBG2,
		"dbg3": DBG3,
		"dbug": DBUG,
		"deco": DECO,
		"decv": DECV,
		"dtac": DTAC,
		"dtde": DTDE,
		"dten": DTEN,
		"dtlv": DTLV,
		"dump": DUMP,
		"else": ELSE,
		"enco": ENCO,
		"fnen": FNEN,
		"fnlv": FNLV,
		"fxen": FXEN,
		"fxlv": FXLV,
		"getv": GETV,
		"ilen": ILEN,
		"illv": ILLV,
		"matc": MATC,
		"mpar": MPAR,
		"paon": PAON,
		"perf": PERF,
		"pllg": PLLG,
		"plod": PLOD,
		"ptck": PTCK,
		"ptcl": PTCL,
		"ptcn": PTCN,
		"ptdi": PTDI,
		"ptds": PTDS,
		"ptha": PTHA,
		"ptmp": PTMP,
		"ptpu": PTPU,
		"ptqu": PTQU,
		"ptre": PTRE,
		"ptrx": PTRX,
		"ptsd": PTSD,
		"ptsp": PTSP,
		"ptst": PTST,
		"pttr": PTTR,
		"ptun": PTUN,
		"rvon": RVON,
		"sdbg": SDBG,
		"setv": SETV,
		"tcen": TCEN,
		"tcfi": TCFI,
		"tclv": TCLV,
		"tcst": TCST,
		"tmrd": TMRD,
		"tmru": TMRU,
		"tmsp": TMSP,
		"tmst": TMST,
		"tmto": TMTO,
		"uact": UACT,
		"ulog": ULOG,
		"vach": VACH,
		"vrsn": VRSN,
		"wait": WAIT,
		"ACDC": ErrACDC,
		"ARGI": ErrARGI,
		"ARGS": ErrARGS,
		"ARRV": ErrARRV,
		"ARTE": ErrARTE,
		"ARUN": ErrARUN,
		"BREF": ErrBREF,
		"CONE": ErrCONE,
		"CONV": ErrCONV,
		"DEAD": ErrDEAD,
		"DIRE": ErrDIRE,
		"DIV0": ErrDIV0,
		"DKEY": ErrDKEY,
		"DOME": ErrDOME,
		"DTDE": ErrDTDE,
		"FILE": ErrFILE,
		"FLOW": ErrFLOW,
		"FOEX": ErrFOEX,
		"ICST": ErrICST,
		"IGRP": ErrIGRP,
		"INOP": ErrINOP,
		"LENE": ErrLENE,
		"LIDE": ErrLIDE,
		"LKUP": ErrLKUP,
		"MALF": ErrMALF,
		"MANY": ErrMANY,
		"MAPE": ErrMAPE,
		"NAME": ErrNAME,
		"NOIM": ErrNOIM,
		"NRTE": ErrNRTE,
		"NSPR": ErrNSPR,
		"NULL": ErrNULL,
		"OMIT": ErrOMIT,
		"OPTM": ErrOPTM,
		"OPTS": ErrOPTS,
		"OPTU": ErrOPTU,
		"OSUF": ErrOSUF,
		"PABV": ErrPABV,
		"PAOV": ErrPAOV,
		"PARE": ErrPARE,
		"PATE": ErrPATE,
		"PLEX": ErrPLEX,
		"PLOD": ErrPLOD,
		"PTFA": ErrPTFA,
		"RNGE": ErrRNGE,
		"SIZE": ErrSIZE,
		"SNAC": ErrSNAC,
		"SNAP": ErrSNAP,
		"SRUN": ErrSRUN,
		"SYSE": ErrSYSE,
		"TIME": ErrTIME,
		"TSTP": ErrTSTP,
		"TYPE": ErrTYPE,
		"UBLK": ErrUBLK,
		"UNAS": ErrUNAS,
		"UNDF": ErrUNDF,
		"UNOP": ErrUNOP,
		"UTF8": ErrUTF8,
		"UTOB": ErrUTOB,
		"VRSN": ErrVRSN,
		"WAIT": ErrWAIT,
		"WCPA": ErrWCPA,
		"WIDE": ErrWIDE,
		"WPAC": ErrWPAC,
	}
)
