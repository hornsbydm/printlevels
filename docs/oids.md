SNMP OIDS
=========

A list of known oids for each model

Generic
-------
``` go
	oid_sys_location string = "1.3.6.1.2.1.1.6.0"
	oid_sys_name     string = "1.3.6.1.2.1.1.5.0"
	oid_sys_contact  string = "1.3.6.1.2.1.1.4.0"
	oid_sys_model    string = "1.3.6.1.2.1.1.1.0"
```

iso.3.6.1.2.1.25.3.5.1.1.1 = Status
3 = idle
4 = printing
5 = warmup

iso.3.6.1.2.1.25.3.5.1.2 = Error State

Bit # | Error
0 | Low Papoer
1 | No PAper
2 | Low Toner
3 | No Toner
4 | door Open
5 | jammed
6 | offline
7 | service request

Konica Color (bizhub )
----------------------

Toner Black,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.4,1.3.6.1.2.1.43.11.1.1.8.1.4,Replacement Model
Toner Cyan,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.1,1.3.6.1.2.1.43.11.1.1.8.1.1,Replacement Model
Toner Magenta,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.2,1.3.6.1.2.1.43.11.1.1.8.1.2,Replacement Model
Toner Gold,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.3,1.3.6.1.2.1.43.11.1.1.8.1.3,Replacement Model
Drum Black,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.11,1.3.6.1.2.1.43.11.1.1.8.1.11,Replacement Model
Drum Cyan,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.5,1.3.6.1.2.1.43.11.1.1.8.1.5,Replacement Model
Drum Magenta,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.7,1.3.6.1.2.1.43.11.1.1.8.1.7,Replacement Model
Drum Gold,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.9,1.3.6.1.2.1.43.11.1.1.8.1.9,Replacement Model
Developer Black,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.12,1.3.6.1.2.1.43.11.1.1.8.1.12,Replacement Model
Developer Cyan,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.6,1.3.6.1.2.1.43.11.1.1.8.1.6,Replacement Model
Developer Magenta,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.8,1.3.6.1.2.1.43.11.1.1.8.1.8,Replacement Model
Developer Gold,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.10,1.3.6.1.2.1.43.11.1.1.8.1.10,Replacement Model
Waste Toner,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.13,1.3.6.1.2.1.43.11.1.1.8.1.13,Replacement Model
Fuser,10.10.2.1,1.3.6.1.2.1.43.11.1.1.9.1.14,1.3.6.1.2.1.43.11.1.1.8.1.14,Replacement Model
Print Meter,10.10.2.1,1.3.6.1.2.1.43.10.2.1.4.1.1,M100000,Replacement Model

Konica Mono (bizhub 284e )
---------------------
iso.3.6.1.2.1.43.11.1.1.6.1.1 = STRING: "Toner (Black)"
iso.3.6.1.2.1.43.11.1.1.6.1.2 = STRING: "Drum Cartridge (Black)"
iso.3.6.1.2.1.43.11.1.1.6.1.3 = STRING: "Developer Cartridge (Black)"
iso.3.6.1.2.1.43.11.1.1.6.1.4 = STRING: "Waste Toner Box"
iso.3.6.1.2.1.43.11.1.1.6.1.5 = STRING: "Fusing Unit"
iso.3.6.1.2.1.43.11.1.1.6.1.6 = STRING: "Image Transfer Belt Unit"
iso.3.6.1.2.1.43.11.1.1.6.1.7 = STRING: "Transfer Roller Unit"
iso.3.6.1.2.1.43.11.1.1.6.1.8 = STRING: "Ozone Filter"
iso.3.6.1.2.1.43.11.1.1.7.1.1 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.7.1.2 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.7.1.3 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.7.1.4 = INTEGER: 1
iso.3.6.1.2.1.43.11.1.1.7.1.5 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.7.1.6 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.7.1.7 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.7.1.8 = INTEGER: 19
iso.3.6.1.2.1.43.11.1.1.8.1.1 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.8.1.2 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.8.1.3 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.8.1.4 = INTEGER: -2
iso.3.6.1.2.1.43.11.1.1.8.1.5 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.8.1.6 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.8.1.7 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.8.1.8 = INTEGER: 100
iso.3.6.1.2.1.43.11.1.1.9.1.1 = INTEGER: 51
iso.3.6.1.2.1.43.11.1.1.9.1.2 = INTEGER: 52
iso.3.6.1.2.1.43.11.1.1.9.1.3 = INTEGER: 91
iso.3.6.1.2.1.43.11.1.1.9.1.4 = INTEGER: -3
iso.3.6.1.2.1.43.11.1.1.9.1.5 = INTEGER: 91
iso.3.6.1.2.1.43.11.1.1.9.1.6 = INTEGER: 81
iso.3.6.1.2.1.43.11.1.1.9.1.7 = INTEGER: 90
iso.3.6.1.2.1.43.11.1.1.9.1.8 = INTEGER: 90

Lexmark M5155
------------



Dell Color
----------

Lexmark XC2235
--------------

iso.3.6.1.2.1.43.11.1.1.6.1.1 = STRING: "Black Cartridge"
iso.3.6.1.2.1.43.11.1.1.6.1.2 = STRING: "Cyan Cartridge"
iso.3.6.1.2.1.43.11.1.1.6.1.3 = STRING: "Transfer Module"
iso.3.6.1.2.1.43.11.1.1.6.1.4 = STRING: "Imaging Kit"
iso.3.6.1.2.1.43.11.1.1.6.1.5 = STRING: "Magenta Cartridge"
iso.3.6.1.2.1.43.11.1.1.6.1.6 = STRING: "Maintenance Kit"
iso.3.6.1.2.1.43.11.1.1.6.1.7 = STRING: "Waste Toner Bottle"
iso.3.6.1.2.1.43.11.1.1.6.1.8 = STRING: "Yellow Cartridge"
iso.3.6.1.2.1.43.11.1.1.7.1.1 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.2 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.3 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.4 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.5 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.6 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.7 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.8 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.8.1.1 = INTEGER: 9000
iso.3.6.1.2.1.43.11.1.1.8.1.2 = INTEGER: 6000
iso.3.6.1.2.1.43.11.1.1.8.1.3 = INTEGER: 125000
iso.3.6.1.2.1.43.11.1.1.8.1.4 = INTEGER: 125000
iso.3.6.1.2.1.43.11.1.1.8.1.5 = INTEGER: 6000
iso.3.6.1.2.1.43.11.1.1.8.1.6 = INTEGER: 125000
iso.3.6.1.2.1.43.11.1.1.8.1.7 = INTEGER: 25000
iso.3.6.1.2.1.43.11.1.1.8.1.8 = INTEGER: 6000
iso.3.6.1.2.1.43.11.1.1.9.1.1 = INTEGER: 5130
iso.3.6.1.2.1.43.11.1.1.9.1.2 = INTEGER: 6000
iso.3.6.1.2.1.43.11.1.1.9.1.3 = INTEGER: 125000
iso.3.6.1.2.1.43.11.1.1.9.1.4 = INTEGER: 125000
iso.3.6.1.2.1.43.11.1.1.9.1.5 = INTEGER: 6000
iso.3.6.1.2.1.43.11.1.1.9.1.6 = INTEGER: 125000
iso.3.6.1.2.1.43.11.1.1.9.1.7 = INTEGER: 15000
iso.3.6.1.2.1.43.11.1.1.9.1.8 = INTEGER: 3240

Toner Black Level	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.1	1.3.6.1.2.1.43.11.1.1.8.1.1	Replacement Model
Imaging Unit Level	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.4	1.3.6.1.2.1.43.11.1.1.8.1.4	Replacement Model
Print Meter	10.10.26.4	1.3.6.1.2.1.43.10.2.1.4.1.1	M100000	Replacement Model
Toner Cyan	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.2	1.3.6.1.2.1.43.11.1.1.8.1.2	Replacement Model
Transfer Module	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.3	1.3.6.1.2.1.43.11.1.1.8.1.3	Replacement Model
Toner Magenta	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.5	1.3.6.1.2.1.43.11.1.1.8.1.5	Replacement Model
Maintenance Kit	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.6	1.3.6.1.2.1.43.11.1.1.8.1.6	Replacement Model
Waste Toner	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.7	1.3.6.1.2.1.43.11.1.1.8.1.7	Replacement Model
Toner Yellow	10.10.26.4	1.3.6.1.2.1.43.11.1.1.9.1.8	1.3.6.1.2.1.43.11.1.1.8.1.8	Replacement Model


Lexmark X748 Color
------------------

iso.3.6.1.2.1.43.11.1.1.6.1.1 = STRING: "Cyan Toner"
iso.3.6.1.2.1.43.11.1.1.6.1.2 = STRING: "Magenta Toner"
iso.3.6.1.2.1.43.11.1.1.6.1.3 = STRING: "Yellow Toner"
iso.3.6.1.2.1.43.11.1.1.6.1.4 = STRING: "Black Toner"
iso.3.6.1.2.1.43.11.1.1.6.1.5 = STRING: "Photo Drum:Cyan"
iso.3.6.1.2.1.43.11.1.1.6.1.6 = STRING: "Photo Drum:Magenta"
iso.3.6.1.2.1.43.11.1.1.6.1.7 = STRING: "Photo Drum:Yellow"
iso.3.6.1.2.1.43.11.1.1.6.1.8 = STRING: "Photo Drum:Black"
iso.3.6.1.2.1.43.11.1.1.6.1.9 = STRING: "Waste Toner Bottle"
iso.3.6.1.2.1.43.11.1.1.6.1.10 = STRING: "Fuser"
iso.3.6.1.2.1.43.11.1.1.6.1.11 = STRING: "Transfer Module"
iso.3.6.1.2.1.43.11.1.1.6.1.12 = STRING: "Separator pad"
iso.3.6.1.2.1.43.11.1.1.7.1.1 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.2 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.3 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.4 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.5 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.6 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.7 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.8 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.9 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.10 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.11 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.7.1.12 = INTEGER: 7
iso.3.6.1.2.1.43.11.1.1.8.1.1 = INTEGER: 10000
iso.3.6.1.2.1.43.11.1.1.8.1.2 = INTEGER: 10000
iso.3.6.1.2.1.43.11.1.1.8.1.3 = INTEGER: 10000
iso.3.6.1.2.1.43.11.1.1.8.1.4 = INTEGER: 12000
iso.3.6.1.2.1.43.11.1.1.8.1.5 = INTEGER: 22000
iso.3.6.1.2.1.43.11.1.1.8.1.6 = INTEGER: 22000
iso.3.6.1.2.1.43.11.1.1.8.1.7 = INTEGER: 22000
iso.3.6.1.2.1.43.11.1.1.8.1.8 = INTEGER: 22000
iso.3.6.1.2.1.43.11.1.1.8.1.9 = INTEGER: 20000
iso.3.6.1.2.1.43.11.1.1.8.1.10 = INTEGER: 120000
iso.3.6.1.2.1.43.11.1.1.8.1.11 = INTEGER: 120000
iso.3.6.1.2.1.43.11.1.1.8.1.12 = INTEGER: 30000
iso.3.6.1.2.1.43.11.1.1.9.1.1 = INTEGER: 9400
iso.3.6.1.2.1.43.11.1.1.9.1.2 = INTEGER: 9600
iso.3.6.1.2.1.43.11.1.1.9.1.3 = INTEGER: 9600
iso.3.6.1.2.1.43.11.1.1.9.1.4 = INTEGER: 11280
iso.3.6.1.2.1.43.11.1.1.9.1.5 = INTEGER: 15620
iso.3.6.1.2.1.43.11.1.1.9.1.6 = INTEGER: 15620
iso.3.6.1.2.1.43.11.1.1.9.1.7 = INTEGER: 15620
iso.3.6.1.2.1.43.11.1.1.9.1.8 = INTEGER: 15620
iso.3.6.1.2.1.43.11.1.1.9.1.9 = INTEGER: 20000
iso.3.6.1.2.1.43.11.1.1.9.1.10 = INTEGER: 92400
iso.3.6.1.2.1.43.11.1.1.9.1.11 = INTEGER: 92400
iso.3.6.1.2.1.43.11.1.1.9.1.12 = INTEGER: -3
