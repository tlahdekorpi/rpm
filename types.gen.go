// Code generated by generate.awk. DO NOT EDIT.

package rpm

const (
	HEADER_I18NTABLE  = 100
	HEADER_REGIONS    = 64
	HEADER_SIGBASE    = 256
	HEADER_IMMUTABLE  = 63
	HEADER_TAGBASE    = 1000
	HEADER_IMAGE      = 61
	HEADER_SIGNATURES = 62
)

//go:generate stringer -output types_string.gen.go -type=TagType
type TagType uint32

const (
	RPMTAG_NAME                        TagType = 1000
	RPMTAG_VERSION                     TagType = 1001
	RPMTAG_RELEASE                     TagType = 1002
	RPMTAG_EPOCH                       TagType = 1003
	RPMTAG_SUMMARY                     TagType = 1004
	RPMTAG_DESCRIPTION                 TagType = 1005
	RPMTAG_BUILDTIME                   TagType = 1006
	RPMTAG_BUILDHOST                   TagType = 1007
	RPMTAG_INSTALLTIME                 TagType = 1008
	RPMTAG_SIZE                        TagType = 1009
	RPMTAG_DISTRIBUTION                TagType = 1010
	RPMTAG_VENDOR                      TagType = 1011
	RPMTAG_GIF                         TagType = 1012
	RPMTAG_XPM                         TagType = 1013
	RPMTAG_LICENSE                     TagType = 1014
	RPMTAG_PACKAGER                    TagType = 1015
	RPMTAG_GROUP                       TagType = 1016
	RPMTAG_CHANGELOG                   TagType = 1017
	RPMTAG_SOURCE                      TagType = 1018
	RPMTAG_PATCH                       TagType = 1019
	RPMTAG_URL                         TagType = 1020
	RPMTAG_OS                          TagType = 1021
	RPMTAG_ARCH                        TagType = 1022
	RPMTAG_PREIN                       TagType = 1023
	RPMTAG_POSTIN                      TagType = 1024
	RPMTAG_PREUN                       TagType = 1025
	RPMTAG_POSTUN                      TagType = 1026
	RPMTAG_OLDFILENAMES                TagType = 1027
	RPMTAG_FILESIZES                   TagType = 1028
	RPMTAG_FILESTATES                  TagType = 1029
	RPMTAG_FILEMODES                   TagType = 1030
	RPMTAG_FILEUIDS                    TagType = 1031
	RPMTAG_FILEGIDS                    TagType = 1032
	RPMTAG_FILERDEVS                   TagType = 1033
	RPMTAG_FILEMTIMES                  TagType = 1034
	RPMTAG_FILEDIGESTS                 TagType = 1035
	RPMTAG_FILELINKTOS                 TagType = 1036
	RPMTAG_FILEFLAGS                   TagType = 1037
	RPMTAG_ROOT                        TagType = 1038
	RPMTAG_FILEUSERNAME                TagType = 1039
	RPMTAG_FILEGROUPNAME               TagType = 1040
	RPMTAG_EXCLUDE                     TagType = 1041
	RPMTAG_EXCLUSIVE                   TagType = 1042
	RPMTAG_ICON                        TagType = 1043
	RPMTAG_SOURCERPM                   TagType = 1044
	RPMTAG_FILEVERIFYFLAGS             TagType = 1045
	RPMTAG_ARCHIVESIZE                 TagType = 1046
	RPMTAG_PROVIDENAME                 TagType = 1047
	RPMTAG_REQUIREFLAGS                TagType = 1048
	RPMTAG_REQUIRENAME                 TagType = 1049
	RPMTAG_REQUIREVERSION              TagType = 1050
	RPMTAG_NOSOURCE                    TagType = 1051
	RPMTAG_NOPATCH                     TagType = 1052
	RPMTAG_CONFLICTFLAGS               TagType = 1053
	RPMTAG_CONFLICTNAME                TagType = 1054
	RPMTAG_CONFLICTVERSION             TagType = 1055
	RPMTAG_DEFAULTPREFIX               TagType = 1056
	RPMTAG_BUILDROOT                   TagType = 1057
	RPMTAG_INSTALLPREFIX               TagType = 1058
	RPMTAG_EXCLUDEARCH                 TagType = 1059
	RPMTAG_EXCLUDEOS                   TagType = 1060
	RPMTAG_EXCLUSIVEARCH               TagType = 1061
	RPMTAG_EXCLUSIVEOS                 TagType = 1062
	RPMTAG_AUTOREQPROV                 TagType = 1063
	RPMTAG_RPMVERSION                  TagType = 1064
	RPMTAG_TRIGGERSCRIPTS              TagType = 1065
	RPMTAG_TRIGGERNAME                 TagType = 1066
	RPMTAG_TRIGGERVERSION              TagType = 1067
	RPMTAG_TRIGGERFLAGS                TagType = 1068
	RPMTAG_TRIGGERINDEX                TagType = 1069
	RPMTAG_VERIFYSCRIPT                TagType = 1079
	RPMTAG_CHANGELOGTIME               TagType = 1080
	RPMTAG_CHANGELOGNAME               TagType = 1081
	RPMTAG_CHANGELOGTEXT               TagType = 1082
	RPMTAG_BROKENMD5                   TagType = 1083
	RPMTAG_PREREQ                      TagType = 1084
	RPMTAG_PREINPROG                   TagType = 1085
	RPMTAG_POSTINPROG                  TagType = 1086
	RPMTAG_PREUNPROG                   TagType = 1087
	RPMTAG_POSTUNPROG                  TagType = 1088
	RPMTAG_BUILDARCHS                  TagType = 1089
	RPMTAG_OBSOLETENAME                TagType = 1090
	RPMTAG_VERIFYSCRIPTPROG            TagType = 1091
	RPMTAG_TRIGGERSCRIPTPROG           TagType = 1092
	RPMTAG_DOCDIR                      TagType = 1093
	RPMTAG_COOKIE                      TagType = 1094
	RPMTAG_FILEDEVICES                 TagType = 1095
	RPMTAG_FILEINODES                  TagType = 1096
	RPMTAG_FILELANGS                   TagType = 1097
	RPMTAG_PREFIXES                    TagType = 1098
	RPMTAG_INSTPREFIXES                TagType = 1099
	RPMTAG_TRIGGERIN                   TagType = 1100
	RPMTAG_TRIGGERUN                   TagType = 1101
	RPMTAG_TRIGGERPOSTUN               TagType = 1102
	RPMTAG_AUTOREQ                     TagType = 1103
	RPMTAG_AUTOPROV                    TagType = 1104
	RPMTAG_CAPABILITY                  TagType = 1105
	RPMTAG_SOURCEPACKAGE               TagType = 1106
	RPMTAG_OLDORIGFILENAMES            TagType = 1107
	RPMTAG_BUILDPREREQ                 TagType = 1108
	RPMTAG_BUILDREQUIRES               TagType = 1109
	RPMTAG_BUILDCONFLICTS              TagType = 1110
	RPMTAG_BUILDMACROS                 TagType = 1111
	RPMTAG_PROVIDEFLAGS                TagType = 1112
	RPMTAG_PROVIDEVERSION              TagType = 1113
	RPMTAG_OBSOLETEFLAGS               TagType = 1114
	RPMTAG_OBSOLETEVERSION             TagType = 1115
	RPMTAG_DIRINDEXES                  TagType = 1116
	RPMTAG_BASENAMES                   TagType = 1117
	RPMTAG_DIRNAMES                    TagType = 1118
	RPMTAG_ORIGDIRINDEXES              TagType = 1119
	RPMTAG_ORIGBASENAMES               TagType = 1120
	RPMTAG_ORIGDIRNAMES                TagType = 1121
	RPMTAG_OPTFLAGS                    TagType = 1122
	RPMTAG_DISTURL                     TagType = 1123
	RPMTAG_PAYLOADFORMAT               TagType = 1124
	RPMTAG_PAYLOADCOMPRESSOR           TagType = 1125
	RPMTAG_PAYLOADFLAGS                TagType = 1126
	RPMTAG_INSTALLCOLOR                TagType = 1127
	RPMTAG_INSTALLTID                  TagType = 1128
	RPMTAG_REMOVETID                   TagType = 1129
	RPMTAG_SHA1RHN                     TagType = 1130
	RPMTAG_RHNPLATFORM                 TagType = 1131
	RPMTAG_PLATFORM                    TagType = 1132
	RPMTAG_PATCHESNAME                 TagType = 1133
	RPMTAG_PATCHESFLAGS                TagType = 1134
	RPMTAG_PATCHESVERSION              TagType = 1135
	RPMTAG_CACHECTIME                  TagType = 1136
	RPMTAG_CACHEPKGPATH                TagType = 1137
	RPMTAG_CACHEPKGSIZE                TagType = 1138
	RPMTAG_CACHEPKGMTIME               TagType = 1139
	RPMTAG_FILECOLORS                  TagType = 1140
	RPMTAG_FILECLASS                   TagType = 1141
	RPMTAG_CLASSDICT                   TagType = 1142
	RPMTAG_FILEDEPENDSX                TagType = 1143
	RPMTAG_FILEDEPENDSN                TagType = 1144
	RPMTAG_DEPENDSDICT                 TagType = 1145
	RPMTAG_SOURCEPKGID                 TagType = 1146
	RPMTAG_FILECONTEXTS                TagType = 1147
	RPMTAG_FSCONTEXTS                  TagType = 1148
	RPMTAG_RECONTEXTS                  TagType = 1149
	RPMTAG_POLICIES                    TagType = 1150
	RPMTAG_PRETRANS                    TagType = 1151
	RPMTAG_POSTTRANS                   TagType = 1152
	RPMTAG_PRETRANSPROG                TagType = 1153
	RPMTAG_POSTTRANSPROG               TagType = 1154
	RPMTAG_DISTTAG                     TagType = 1155
	RPMTAG_OLDSUGGESTSNAME             TagType = 1156
	RPMTAG_OLDSUGGESTSVERSION          TagType = 1157
	RPMTAG_OLDSUGGESTSFLAGS            TagType = 1158
	RPMTAG_OLDENHANCESNAME             TagType = 1159
	RPMTAG_OLDENHANCESVERSION          TagType = 1160
	RPMTAG_OLDENHANCESFLAGS            TagType = 1161
	RPMTAG_PRIORITY                    TagType = 1162
	RPMTAG_CVSID                       TagType = 1163
	RPMTAG_BLINKPKGID                  TagType = 1164
	RPMTAG_BLINKHDRID                  TagType = 1165
	RPMTAG_BLINKNEVRA                  TagType = 1166
	RPMTAG_FLINKPKGID                  TagType = 1167
	RPMTAG_FLINKHDRID                  TagType = 1168
	RPMTAG_FLINKNEVRA                  TagType = 1169
	RPMTAG_PACKAGEORIGIN               TagType = 1170
	RPMTAG_TRIGGERPREIN                TagType = 1171
	RPMTAG_BUILDSUGGESTS               TagType = 1172
	RPMTAG_BUILDENHANCES               TagType = 1173
	RPMTAG_SCRIPTSTATES                TagType = 1174
	RPMTAG_SCRIPTMETRICS               TagType = 1175
	RPMTAG_BUILDCPUCLOCK               TagType = 1176
	RPMTAG_FILEDIGESTALGOS             TagType = 1177
	RPMTAG_VARIANTS                    TagType = 1178
	RPMTAG_XMAJOR                      TagType = 1179
	RPMTAG_XMINOR                      TagType = 1180
	RPMTAG_REPOTAG                     TagType = 1181
	RPMTAG_KEYWORDS                    TagType = 1182
	RPMTAG_BUILDPLATFORMS              TagType = 1183
	RPMTAG_PACKAGECOLOR                TagType = 1184
	RPMTAG_PACKAGEPREFCOLOR            TagType = 1185
	RPMTAG_XATTRSDICT                  TagType = 1186
	RPMTAG_FILEXATTRSX                 TagType = 1187
	RPMTAG_DEPATTRSDICT                TagType = 1188
	RPMTAG_CONFLICTATTRSX              TagType = 1189
	RPMTAG_OBSOLETEATTRSX              TagType = 1190
	RPMTAG_PROVIDEATTRSX               TagType = 1191
	RPMTAG_REQUIREATTRSX               TagType = 1192
	RPMTAG_BUILDPROVIDES               TagType = 1193
	RPMTAG_BUILDOBSOLETES              TagType = 1194
	RPMTAG_DBINSTANCE                  TagType = 1195
	RPMTAG_NVRA                        TagType = 1196
	RPMTAG_FILENAMES                   TagType = 5000
	RPMTAG_FILEPROVIDE                 TagType = 5001
	RPMTAG_FILEREQUIRE                 TagType = 5002
	RPMTAG_FSNAMES                     TagType = 5003
	RPMTAG_FSSIZES                     TagType = 5004
	RPMTAG_TRIGGERCONDS                TagType = 5005
	RPMTAG_TRIGGERTYPE                 TagType = 5006
	RPMTAG_ORIGFILENAMES               TagType = 5007
	RPMTAG_LONGFILESIZES               TagType = 5008
	RPMTAG_LONGSIZE                    TagType = 5009
	RPMTAG_FILECAPS                    TagType = 5010
	RPMTAG_FILEDIGESTALGO              TagType = 5011
	RPMTAG_BUGURL                      TagType = 5012
	RPMTAG_EVR                         TagType = 5013
	RPMTAG_NVR                         TagType = 5014
	RPMTAG_NEVR                        TagType = 5015
	RPMTAG_NEVRA                       TagType = 5016
	RPMTAG_HEADERCOLOR                 TagType = 5017
	RPMTAG_VERBOSE                     TagType = 5018
	RPMTAG_EPOCHNUM                    TagType = 5019
	RPMTAG_PREINFLAGS                  TagType = 5020
	RPMTAG_POSTINFLAGS                 TagType = 5021
	RPMTAG_PREUNFLAGS                  TagType = 5022
	RPMTAG_POSTUNFLAGS                 TagType = 5023
	RPMTAG_PRETRANSFLAGS               TagType = 5024
	RPMTAG_POSTTRANSFLAGS              TagType = 5025
	RPMTAG_VERIFYSCRIPTFLAGS           TagType = 5026
	RPMTAG_TRIGGERSCRIPTFLAGS          TagType = 5027
	RPMTAG_COLLECTIONS                 TagType = 5029
	RPMTAG_POLICYNAMES                 TagType = 5030
	RPMTAG_POLICYTYPES                 TagType = 5031
	RPMTAG_POLICYTYPESINDEXES          TagType = 5032
	RPMTAG_POLICYFLAGS                 TagType = 5033
	RPMTAG_VCS                         TagType = 5034
	RPMTAG_ORDERNAME                   TagType = 5035
	RPMTAG_ORDERVERSION                TagType = 5036
	RPMTAG_ORDERFLAGS                  TagType = 5037
	RPMTAG_MSSFMANIFEST                TagType = 5038
	RPMTAG_MSSFDOMAIN                  TagType = 5039
	RPMTAG_INSTFILENAMES               TagType = 5040
	RPMTAG_REQUIRENEVRS                TagType = 5041
	RPMTAG_PROVIDENEVRS                TagType = 5042
	RPMTAG_OBSOLETENEVRS               TagType = 5043
	RPMTAG_CONFLICTNEVRS               TagType = 5044
	RPMTAG_FILENLINKS                  TagType = 5045
	RPMTAG_RECOMMENDNAME               TagType = 5046
	RPMTAG_RECOMMENDVERSION            TagType = 5047
	RPMTAG_RECOMMENDFLAGS              TagType = 5048
	RPMTAG_SUGGESTNAME                 TagType = 5049
	RPMTAG_SUGGESTVERSION              TagType = 5050
	RPMTAG_SUGGESTFLAGS                TagType = 5051
	RPMTAG_SUPPLEMENTNAME              TagType = 5052
	RPMTAG_SUPPLEMENTVERSION           TagType = 5053
	RPMTAG_SUPPLEMENTFLAGS             TagType = 5054
	RPMTAG_ENHANCENAME                 TagType = 5055
	RPMTAG_ENHANCEVERSION              TagType = 5056
	RPMTAG_ENHANCEFLAGS                TagType = 5057
	RPMTAG_RECOMMENDNEVRS              TagType = 5058
	RPMTAG_SUGGESTNEVRS                TagType = 5059
	RPMTAG_SUPPLEMENTNEVRS             TagType = 5060
	RPMTAG_ENHANCENEVRS                TagType = 5061
	RPMTAG_ENCODING                    TagType = 5062
	RPMTAG_FILETRIGGERIN               TagType = 5063
	RPMTAG_FILETRIGGERUN               TagType = 5064
	RPMTAG_FILETRIGGERPOSTUN           TagType = 5065
	RPMTAG_FILETRIGGERSCRIPTS          TagType = 5066
	RPMTAG_FILETRIGGERSCRIPTPROG       TagType = 5067
	RPMTAG_FILETRIGGERSCRIPTFLAGS      TagType = 5068
	RPMTAG_FILETRIGGERNAME             TagType = 5069
	RPMTAG_FILETRIGGERINDEX            TagType = 5070
	RPMTAG_FILETRIGGERVERSION          TagType = 5071
	RPMTAG_FILETRIGGERFLAGS            TagType = 5072
	RPMTAG_TRANSFILETRIGGERIN          TagType = 5073
	RPMTAG_TRANSFILETRIGGERUN          TagType = 5074
	RPMTAG_TRANSFILETRIGGERPOSTUN      TagType = 5075
	RPMTAG_TRANSFILETRIGGERSCRIPTS     TagType = 5076
	RPMTAG_TRANSFILETRIGGERSCRIPTPROG  TagType = 5077
	RPMTAG_TRANSFILETRIGGERSCRIPTFLAGS TagType = 5078
	RPMTAG_TRANSFILETRIGGERNAME        TagType = 5079
	RPMTAG_TRANSFILETRIGGERINDEX       TagType = 5080
	RPMTAG_TRANSFILETRIGGERVERSION     TagType = 5081
	RPMTAG_TRANSFILETRIGGERFLAGS       TagType = 5082
	RPMTAG_REMOVEPATHPOSTFIXES         TagType = 5083
	RPMTAG_FILETRIGGERPRIORITIES       TagType = 5084
	RPMTAG_TRANSFILETRIGGERPRIORITIES  TagType = 5085
	RPMTAG_FILETRIGGERCONDS            TagType = 5086
	RPMTAG_FILETRIGGERTYPE             TagType = 5087
	RPMTAG_TRANSFILETRIGGERCONDS       TagType = 5088
	RPMTAG_TRANSFILETRIGGERTYPE        TagType = 5089
	RPMTAG_FILESIGNATURES              TagType = 5090
	RPMTAG_FILESIGNATURELENGTH         TagType = 5091
	RPMTAG_PAYLOADDIGEST               TagType = 5092
	RPMTAG_PAYLOADDIGESTALGO           TagType = 5093
	RPMTAG_AUTOINSTALLED               TagType = 5094
	RPMTAG_IDENTITY                    TagType = 5095
	RPMTAG_MODULARITYLABEL             TagType = 5096
	RPMTAG_PAYLOADDIGESTALT            TagType = 5097
	RPMTAG_HEADERI18NTABLE             TagType = HEADER_I18NTABLE
	RPMTAG_HEADERIMAGE                 TagType = HEADER_IMAGE
	RPMTAG_HEADERIMMUTABLE             TagType = HEADER_IMMUTABLE
	RPMTAG_HEADERREGIONS               TagType = HEADER_REGIONS
	RPMTAG_SIG_BASE                    TagType = HEADER_SIGBASE
	RPMTAG_HEADERSIGNATURES            TagType = HEADER_SIGNATURES
	RPMTAG_SIGSIZE                     TagType = RPMTAG_SIG_BASE + 1
	RPMTAG_PUBKEYS                     TagType = RPMTAG_SIG_BASE + 10
	RPMTAG_DSAHEADER                   TagType = RPMTAG_SIG_BASE + 11
	RPMTAG_RSAHEADER                   TagType = RPMTAG_SIG_BASE + 12
	RPMTAG_SHA1HEADER                  TagType = RPMTAG_SIG_BASE + 13
	RPMTAG_LONGSIGSIZE                 TagType = RPMTAG_SIG_BASE + 14
	RPMTAG_LONGARCHIVESIZE             TagType = RPMTAG_SIG_BASE + 15
	RPMTAG_SHA256HEADER                TagType = RPMTAG_SIG_BASE + 17
	RPMTAG_SIGLEMD5_1                  TagType = RPMTAG_SIG_BASE + 2
	RPMTAG_SIGPGP                      TagType = RPMTAG_SIG_BASE + 3
	RPMTAG_SIGLEMD5_2                  TagType = RPMTAG_SIG_BASE + 4
	RPMTAG_SIGMD5                      TagType = RPMTAG_SIG_BASE + 5
	RPMTAG_SIGGPG                      TagType = RPMTAG_SIG_BASE + 6
	RPMTAG_SIGPGP5                     TagType = RPMTAG_SIG_BASE + 7
	RPMTAG_BADSHA1_1                   TagType = RPMTAG_SIG_BASE + 8
	RPMTAG_BADSHA1_2                   TagType = RPMTAG_SIG_BASE + 9
	RPMTAG_C                           TagType = RPMTAG_CONFLICTNAME
	RPMTAG_SVNID                       TagType = RPMTAG_CVSID
	RPMTAG_ENHANCES                    TagType = RPMTAG_ENHANCENAME
	RPMTAG_E                           TagType = RPMTAG_EPOCH
	RPMTAG_FILEMD5S                    TagType = RPMTAG_FILEDIGESTS
	RPMTAG_N                           TagType = RPMTAG_NAME
	RPMTAG_O                           TagType = RPMTAG_OBSOLETENAME
	RPMTAG_OLDENHANCES                 TagType = RPMTAG_OLDENHANCESNAME
	RPMTAG_OLDSUGGESTS                 TagType = RPMTAG_OLDSUGGESTSNAME
	RPMTAG_P                           TagType = RPMTAG_PROVIDENAME
	RPMTAG_RECOMMENDS                  TagType = RPMTAG_RECOMMENDNAME
	RPMTAG_R                           TagType = RPMTAG_RELEASE
	RPMTAG_REQUIRES                    TagType = RPMTAG_REQUIRENAME
	RPMTAG_HDRID                       TagType = RPMTAG_SHA1HEADER
	RPMTAG_PKGID                       TagType = RPMTAG_SIGMD5
	RPMTAG_SUGGESTS                    TagType = RPMTAG_SUGGESTNAME
	RPMTAG_SUPPLEMENTS                 TagType = RPMTAG_SUPPLEMENTNAME
	RPMTAG_V                           TagType = RPMTAG_VERSION
)

type SigTagType = TagType

const (
	RPMSIGTAG_SIZE                SigTagType = 1000
	RPMSIGTAG_LEMD5_1             SigTagType = 1001
	RPMSIGTAG_PGP                 SigTagType = 1002
	RPMSIGTAG_LEMD5_2             SigTagType = 1003
	RPMSIGTAG_MD5                 SigTagType = 1004
	RPMSIGTAG_GPG                 SigTagType = 1005
	RPMSIGTAG_PGP5                SigTagType = 1006
	RPMSIGTAG_PAYLOADSIZE         SigTagType = 1007
	RPMSIGTAG_RESERVEDSPACE       SigTagType = 1008
	RPMSIGTAG_BADSHA1_1           SigTagType = RPMTAG_BADSHA1_1
	RPMSIGTAG_BADSHA1_2           SigTagType = RPMTAG_BADSHA1_2
	RPMSIGTAG_DSA                 SigTagType = RPMTAG_DSAHEADER
	RPMSIGTAG_LONGARCHIVESIZE     SigTagType = RPMTAG_LONGARCHIVESIZE
	RPMSIGTAG_LONGSIZE            SigTagType = RPMTAG_LONGSIGSIZE
	RPMSIGTAG_RSA                 SigTagType = RPMTAG_RSAHEADER
	RPMSIGTAG_SHA1                SigTagType = RPMTAG_SHA1HEADER
	RPMSIGTAG_SHA256              SigTagType = RPMTAG_SHA256HEADER
	RPMSIGTAG_FILESIGNATURES      SigTagType = RPMTAG_SIG_BASE + 18
	RPMSIGTAG_FILESIGNATURELENGTH SigTagType = RPMTAG_SIG_BASE + 19
)

var sigTagString = map[TagType]string{
	RPMSIGTAG_SIZE:                "RPMSIGTAG_SIZE",
	RPMSIGTAG_LEMD5_1:             "RPMSIGTAG_LEMD5_1",
	RPMSIGTAG_PGP:                 "RPMSIGTAG_PGP",
	RPMSIGTAG_LEMD5_2:             "RPMSIGTAG_LEMD5_2",
	RPMSIGTAG_MD5:                 "RPMSIGTAG_MD5",
	RPMSIGTAG_GPG:                 "RPMSIGTAG_GPG",
	RPMSIGTAG_PGP5:                "RPMSIGTAG_PGP5",
	RPMSIGTAG_PAYLOADSIZE:         "RPMSIGTAG_PAYLOADSIZE",
	RPMSIGTAG_RESERVEDSPACE:       "RPMSIGTAG_RESERVEDSPACE",
	RPMSIGTAG_BADSHA1_1:           "RPMSIGTAG_BADSHA1_1",
	RPMSIGTAG_BADSHA1_2:           "RPMSIGTAG_BADSHA1_2",
	RPMSIGTAG_DSA:                 "RPMSIGTAG_DSA",
	RPMSIGTAG_LONGARCHIVESIZE:     "RPMSIGTAG_LONGARCHIVESIZE",
	RPMSIGTAG_LONGSIZE:            "RPMSIGTAG_LONGSIZE",
	RPMSIGTAG_RSA:                 "RPMSIGTAG_RSA",
	RPMSIGTAG_SHA1:                "RPMSIGTAG_SHA1",
	RPMSIGTAG_SHA256:              "RPMSIGTAG_SHA256",
	RPMSIGTAG_FILESIGNATURES:      "RPMSIGTAG_FILESIGNATURES",
	RPMSIGTAG_FILESIGNATURELENGTH: "RPMSIGTAG_FILESIGNATURELENGTH",
}

const (
	RPM_NULL_TYPE         = 0
	RPM_CHAR_TYPE         = 1
	RPM_INT8_TYPE         = 2
	RPM_INT16_TYPE        = 3
	RPM_INT32_TYPE        = 4
	RPM_INT64_TYPE        = 5
	RPM_STRING_TYPE       = 6
	RPM_BIN_TYPE          = 7
	RPM_STRING_ARRAY_TYPE = 8
	RPM_I18NSTRING_TYPE   = 9
	RPM_MIN_TYPE          = 0
	RPM_MASK_TYPE         = 0x0000ffff
	RPM_FORCEFREE_TYPE    = 0xff
	RPM_MAX_TYPE          = 9
)
const (
	RPMSENSE_LESS          = (1 << 1)
	RPMSENSE_SCRIPT_POST   = (1 << 10)
	RPMSENSE_SCRIPT_PREUN  = (1 << 11)
	RPMSENSE_SCRIPT_POSTUN = (1 << 12)
	RPMSENSE_SCRIPT_VERIFY = (1 << 13)
	RPMSENSE_FIND_REQUIRES = (1 << 14)
	RPMSENSE_FIND_PROVIDES = (1 << 15)
	RPMSENSE_TRIGGERIN     = (1 << 16)
	RPMSENSE_TRIGGERUN     = (1 << 17)
	RPMSENSE_TRIGGERPOSTUN = (1 << 18)
	RPMSENSE_MISSINGOK     = (1 << 19)
	RPMSENSE_GREATER       = (1 << 2)
	RPMSENSE_RPMLIB        = (1 << 24)
	RPMSENSE_TRIGGERPREIN  = (1 << 25)
	RPMSENSE_KEYRING       = (1 << 26)
	RPMSENSE_CONFIG        = (1 << 28)
	RPMSENSE_META          = (1 << 29)
	RPMSENSE_EQUAL         = (1 << 3)
	RPMSENSE_POSTTRANS     = (1 << 5)
	RPMSENSE_PREREQ        = (1 << 6)
	RPMSENSE_PRETRANS      = (1 << 7)
	RPMSENSE_INTERP        = (1 << 8)
	RPMSENSE_SCRIPT_PRE    = (1 << 9)
	RPMSENSE_ANY           = 0
)
const (
	RPMFILE_CONFIG    = (1 << 0)
	RPMFILE_DOC       = (1 << 1)
	RPMFILE_ICON      = (1 << 2)
	RPMFILE_MISSINGOK = (1 << 3)
	RPMFILE_NOREPLACE = (1 << 4)
	RPMFILE_SPECFILE  = (1 << 5)
	RPMFILE_GHOST     = (1 << 6)
	RPMFILE_LICENSE   = (1 << 7)
	RPMFILE_README    = (1 << 8)
	RPMFILE_PUBKEY    = (1 << 11)
	RPMFILE_ARTIFACT  = (1 << 12)
	RPMFILE_NONE      = 0
)
const (
	RPMVERIFY_FILEDIGEST      = (1 << 0)
	RPMVERIFY_FILESIZE        = (1 << 1)
	RPMVERIFY_CONTEXTS        = (1 << 15)
	RPMVERIFY_LINKTO          = (1 << 2)
	RPMVERIFY_READLINKFAIL    = (1 << 28)
	RPMVERIFY_READFAIL        = (1 << 29)
	RPMVERIFY_USER            = (1 << 3)
	RPMVERIFY_LSTATFAIL       = (1 << 30)
	RPMVERIFY_LGETFILECONFAIL = (1 << 31)
	RPMVERIFY_GROUP           = (1 << 4)
	RPMVERIFY_MTIME           = (1 << 5)
	RPMVERIFY_MODE            = (1 << 6)
	RPMVERIFY_RDEV            = (1 << 7)
	RPMVERIFY_CAPS            = (1 << 8)
	RPMVERIFY_NONE            = 0
)
const (
	PGPHASHALGO_MD5         = 1
	PGPHASHALGO_SHA1        = 2
	PGPHASHALGO_RIPEMD160   = 3
	PGPHASHALGO_MD2         = 5
	PGPHASHALGO_TIGER192    = 6
	PGPHASHALGO_HAVAL_5_160 = 7
	PGPHASHALGO_SHA256      = 8
	PGPHASHALGO_SHA384      = 9
	PGPHASHALGO_SHA512      = 10
	PGPHASHALGO_SHA224      = 11
)