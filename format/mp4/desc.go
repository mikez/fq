package mp4

import "github.com/wader/fq/pkg/scalar"

// based on https://github.com/HexFiend/HexFiend/blob/master/templates/Media/MOV.tcl

var boxDescriptions = scalar.StrToDescription{
	"dimg": "Derived image",
	"cdsc": "Content description",
	"ainf": "Asset information to identify, license and play",
	"albm": "Album title and track number for media",
	"alou": "Album loudness base",
	"angl": "Name of the camera angle through which the clip was shot",
	"assp": "Alternative startup sequence properties",
	"auth": "Author of the media",
	"auxC": "Image properties for auxiliary images clap clean aperture",
	"avcn": "AVC NAL Unit Storage Box",
	"bidx": "Box Index",
	"bloc": "Base location and purchase location for license acquisition",
	"bpcc": "Bits per component",
	"btrt": "Bitrate",
	"buff": "Buffering information",
	"bxml": "Binary XML container",
	"ccid": "OMA DRM Content ID",
	"cdef": "Type and ordering of the components within the codestream",
	"cinf": "Complete track information",
	"clfn": "Name of the clip file",
	"clid": "Identifier of the clip",
	"clip": "Visual clipping region container",
	"clsf": "Classification of the media",
	"cmap": "Mapping between a palette and codestream components",
	"cmid": "Identifier of the camera",
	"cmnm": "Name that identifies the camera",
	"co64": "64-bit chunk offset",
	"coin": "Content Information Box",
	"coll": "Name of the collection from which the media comes",
	"colr": "Specifies the colourspace of the image",
	"cprt": "Copyright etc",
	"crgn": "Visual clipping region definition",
	"crhd": "Reserved for ClockReferenceStream header",
	"csgp": "Compact sample to group",
	"cslg": "Composition to decode timeline mapping",
	"ctab": "Track color-table",
	"ctts": "Composition time to sample",
	"cvru": "OMA DRM Cover URI",
	"dac4": "Dolby AC-4 stream descriptor",
	"date": "Date and time, formatted according to ISO 8601, when the content was created. For clips captured by recording devices, this is typically the date and time when the clip’s recording started",
	"dcfD": "Marlin DCF Duration, user-data atom type",
	"dec3": "E-AC-3 (Dolby Digital Plus) stream descriptor",
	"dihd": "Data Integrity Hash",
	"dinf": "Data information box, container",
	"dint": "Data Integrity",
	"dref": "Data reference box, declares source(s) of media data in track",
	"dscp": "Media description",
	"dsgd": "DVB Sample Group Description Box",
	"dstg": "DVB Sample to Group Box",
	"edts": "Edit list container",
	"elng": "Extended Language Tag",
	"elst": "An edit list",
	"emsg": "Event message",
	"esds": "Elementary stream descriptor",
	"etyp": "Extended type and type combination",
	"evti": "Event information",
	"fdel": "File delivery information (item info extension)",
	"feci": "FEC Informatiom",
	"fecr": "FEC Reservoir",
	"fidx": "Box File Index",
	"fiel": "Video field order",
	"fiin": "FD Item Information",
	"fire": "File Reservoir",
	"fpar": "File Partition",
	"free": "Free space",
	"frma": "Original format box",
	"frpa": "Front Part",
	"ftyp": "File type and compatibility",
	"gitn": "Group ID to name",
	"gnre": "Media genre",
	"grpi": "OMA DRM Group ID",
	"grpl": "Groups List box",
	"hdlr": "Handler, declares the media (handler) type",
	"hinf": "Hint information",
	"hmhd": "Hint media header, overall information (hint track only)",
	"hnti": "Hint information",
	"hpix": "Hipix Rich Picture (user-data or meta-data)",
	"icnu": "OMA DRM Icon URI",
	"ID32": "ID3 version 2 container",
	"idat": "Item data",
	"ihdr": "Image Header",
	"iinf": "Item information",
	"iloc": "Item location",
	"imap": "Track input map definition",
	"imda": "Identified media data",
	"imif": "IPMP Information box",
	"infe": "Item information entry",
	"infu": "OMA DRM Info URL",
	"iods": "Object Descriptor container box",
	"ipco": "ItemPropertyContainerBox",
	"iphd": "Reserved for IPMP Stream header",
	"ipma": "ItemPropertyAssociation",
	"ipmc": "IPMP Control Box",
	"ipro": "Item protection",
	"iprp": "Item Properties Box",
	"iref": "Item reference",
	"irot": "Image rotation imir image mirroring",
	"ispe": "Image spatial extents",
	"j2kH": "JPEG 2000 header item property",
	"jP  ": "JPEG 2000 Signature",
	"jp2c": "JPEG 2000 contiguous codestream",
	"jp2h": "Header",
	"jp2i": "Intellectual property information",
	"kind": "Track kind",
	"kmat": "Compressed visual track matte",
	"kywd": "Media keywords",
	"leva": "Leval assignment",
	"load": "Track pre-load definitions",
	"loci": "Media location information",
	"loop": "Looping behavior",
	"lrcu": "OMA DRM Lyrics URI",
	"ludt": "Track loudness container",
	"m7hd": "Reserved for MPEG7Stream header",
	"manu": "Manufacturer name of the camera",
	"matt": "Visual track matte for compositing",
	"md5i": "MD5IntegrityBox",
	"mdat": "Media data container",
	"mdhd": "Media header, overall information about the media",
	"mdia": "Container for the media information in a track",
	"mdri": "Mutable DRM information",
	"meco": "Additional metadata container",
	"mehd": "Movie extends header box",
	"mere": "Metabox relation",
	"meta": "Metadata container",
	"mfhd": "Movie fragment header",
	"mfra": "Movie fragment random access",
	"mfro": "Movie fragment random access offset",
	"minf": "Media information container",
	"mjhd": "Reserved for MPEG-J Stream header",
	"modl": "Model name of the camera",
	"moof": "Movie fragment",
	"moov": "Container for all the meta-data",
	"mstv": "MVC sub track view box",
	"mvcg": "Multiview group",
	"mvci": "Multiview Information",
	"mvdr": "MVDDepthResolutionBox",
	"mvex": "Movie extends box",
	"mvhd": "Movie header, overall declarations",
	"mvra": "Multiview Relation Attribute",
	"nmhd": "Null media header, overall information (some tracks only)",
	"ochd": "Reserved for ObjectContentInfoStream header",
	"odaf": "OMA DRM Access Unit Format",
	"odda": "OMA DRM Content Object",
	"odhd": "Reserved for ObjectDescriptorStream header",
	"odhe": "OMA DRM Discrete Media Headers",
	"odrb": "OMA DRM Rights Object",
	"odrm": "OMA DRM Container",
	"odtt": "OMA DRM Transaction Tracking",
	"ohdr": "OMA DRM Common headers",
	"orie": "Orientation information",
	"padb": "Sample padding bits",
	"paen": "Partition Entry",
	"pasp": "Pixel aspect ratio",
	"pclr": "Palette which maps a single component in index space to a multiple- component image",
	"pdat": "Partial Data",
	"pdin": "Progressive download information",
	"perf": "Media performer name",
	"pfhd": "Partial File Header",
	"pfil": "Partial File",
	"pitm": "Primary item reference",
	"pixi": "Pixel information",
	"ploc": "Partial Segment Location",
	"pnot": "Preview container",
	"prft": "Producer reference time",
	"pseg": "Partial Segment",
	"pshd": "Partial Segment Header",
	"pssh": "Protection system specific header",
	"ptle": "Partial Top Level Entry",
	"reel": "Name of the tape reel",
	"res ": "Grid resolution",
	"resc": "Grid resolution at which the image was captured",
	"resd": "Default grid resolution at which the image should be displayed",
	"rinf": "Restricted scheme information box",
	"rloc": "Relative location",
	"rtng": "Media rating",
	"saio": "Sample auxiliary information offsets",
	"saiz": "Sample auxiliary information sizes",
	"sbgp": "Sample to Group box",
	"scen": "Name of the scene for which the clip was shot",
	"schi": "Scheme information box",
	"schm": "Scheme type box",
	"sdep": "Sample dependency",
	"sdhd": "Reserved for SceneDescriptionStream header",
	"sdtp": "Sample Dependency Flags",
	"sdvp": "SD Profile Box",
	"segr": "File delivery session group",
	"seii": "SEI information box",
	"senc": "Sample specific encryption data",
	"sgpd": "Sample group definition box",
	"shot": "Name that identifies the shot",
	"sidx": "Segment Index Box",
	"sinf": "Protection scheme information box",
	"skip": "Free space",
	"slno": "Serial number of the camera",
	"smhd": "Sound media header, overall information (sound track only)",
	"srmb": "System Renewability Message",
	"srmc": "System Renewability Message container",
	"srpp": "STRP Process",
	"ssix": "Sub-sample index",
	"sstl": "SVC sub track layer box",
	"stbl": "Sample table box, container for the time/space map",
	"stco": "Chunk offset, partial data-offset information",
	"stdp": "Sample degradation priority",
	"sthd": "Subtitle Media Header Box",
	"stmg": "MVC sub track multiview group box",
	"strd": "Sub-track definition",
	"stri": "Sub-track information",
	"strk": "Sub track information",
	"stsc": "Sample-to-chunk, partial data-offset information",
	"stsd": "Sample descriptions (codec types, initialization etc.)",
	"stsg": "Sub-track sample grouping",
	"stsh": "Shadow sync sample table",
	"stss": "Sync sample table (random access points)",
	"stsz": "Sample sizes (framing)",
	"stti": "Sub track tier box",
	"stts": "Sample time-to-sample",
	"styp": "Segment Type Box",
	"stz2": "Compact sample sizes (framing)",
	"subs": "Sub-sample information",
	"surl": "Source URL",
	"swtc": "Multiview Group Relation",
	"tenc": "Track Encryption",
	"tfad": "Track fragment adjustment box",
	"tfdt": "Track fragment decode time",
	"tfhd": "Track fragment header",
	"tfma": "Track fragment media adjustment box",
	"tfra": "Track fragment radom access",
	"thmb": "Thumbnail image of the media",
	"tibr": "Tier Bit rate",
	"tiri": "Tier Information",
	"titl": "Media title",
	"tkhd": "Track header, overall information about the track",
	"tlou": "Track loudness base",
	"traf": "Track fragment",
	"trak": "Container for an individual track or stream",
	"tref": "Track reference container",
	"trep": "Track extension properties",
	"trex": "Track extends defaults",
	"trgr": "Track grouping information",
	"trik": "Facilitates random access and trick play modes",
	"trun": "Track fragment run",
	"tsel": "Track selection",
	"tstb": "TileSubTrackGroupBox",
	"ttyp": "Track type and compatibility",
	"tyco": "Type and-combination",
	"udta": "User-data",
	"uinf": "A tool by which a vendor may provide access to additional information associated with a UUID",
	"UITS": "Unique Identifier Technology Solution",
	"ulst": "A list of UUID’s",
	"urat": "User 'star' rating of the media",
	"url ": "A URL",
	"uuid": "User-extension box",
	"vmhd": "Video media header, overall information (video track only)",
	"vwdi": "Multiview Scene Information",
	"wide": "Expansion space reservation",
	"xml ": "XML container",
	"yrrc": "Year when media was recorded",

	// from https://wiki.multimedia.cx/index.php/FFmpeg_Metadata
	"©nam": "Title",
	"©ART": "Author",
	"aART": "Album_artist",
	"©alb": "Album",
	"©grp": "Grouping",
	"©wrt": "Composer",
	"©day": "Year",
	"trkn": "Track",
	"©cmt": "Comment",
	"©gen": "Genre",
	"©cpy": "Copyright",
	"desc": "Description",
	"ldes": "Synopsis",
	"tvsh": "Show",
	"tven": "Episode_id",
	"tvnn": "Network",
	"©lyr": "Lyrics",

	"©too": "Encoder",

	"\x00\x00\x00\x00": "Terminator Atom",
}
