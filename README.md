# Quartiles

Find quartiles ([apple news+ new game](https://www.apple.com/newsroom/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/))

You can make a screenshot and use iphone text extraction to copy paste the list of fragments (or type it)

## Installation

From source
```
go install github.com/ldemailly/quartiles@latest
```

Or see the numerous binaries in https://github.com/ldemailly/quartiles/releases

Or docker `fortio/quartiles:latest`

Or brew `brew install fortio/tap/quartiles`

(I manage the fortio org and usually put everything there but this one is completely unrelated outside of it's usage of https://pkg.go.dev/fortio.org/sets#Tuplets so for now it is hosted here in `ldemailly` yet uses fortio's org for docker and brew)

## Example

Make a screenshot, select the word fragments, paste to textedit, convert to plain text, paste to `quartiles`

Example:
```
$ docker run -ti fortio/quartiles:latest "nder
nt
sti
bi
lo
thu
al
con
ies
VO
mon
ba
tue
rd
fri
lit
SW
ck
wed
gam
"
```

```
16:49:38.367 [INF] Words in dictionary : 132735
16:49:38.367 [INF] Looking at fragments:
lb	mac	ash	pa
tri	ns	ar	arch
ect	fl	el	oo
bu	ora	al	tes
pre	dic	tab	le
16:49:38.367 [INF] Checking 2 parts combinations:
al-ar : alar
bu-lb : bulb
bu-ns : buns
el-ect : elect
fl-ash : flash
fl-ora : flora
le-al : leal
le-ash : leash
le-ns : lens
mac-le : macle
ora-tes : orates
pa-le : pale
pa-ns : pans
pa-tes : pates
tab-le : table
tab-oo : taboo
tri-al : trial
16:49:38.367 [INF] Checking 3 parts combinations:
bu-lb-ar : bulbar
le-ar-ns : learns
pa-tri-arch : patriarch
pre-tri-al : pretrial
tes-tab-le : testable
tri-bu-tes : tributes
16:49:38.370 [INF] Checking 4 parts combinations:
el-ect-ora-tes : electorates
fl-ash-bu-lb : flashbulb
mac-ar-oo-ns : macaroons
pa-tri-arch-al : patriarchal
pre-dic-tab-le : predictable
```
## Usage

```
quartiles 0.2.0 usage:
	quartiles [flags] word fragments...
Finds all words in dictionary that can be made from the given fragments
or 1 of the special arguments
	quartiles {help|envhelp|version|buildinfo}
flags:
  -4	Only find the 4-part words
  -dict path
    	Dictionary file path to use, instead of default embedded one
```

## What is the Quartiles game?

https://www.apple.com/newsroom/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/


<img src="https://github.com/ldemailly/quartiles/assets/3664595/4ef37a1b-86e8-4841-88fa-c36bc5e5838b" width="50%">

## Wordlist

The `words` used in Dockerfile is from Ubuntu's
`apt install wamerican-large` american-english-large words-large
removing 's and all caps acronyms
```
grep -v "'" /usr/share/dict/american-english-large | grep -v -E "^[A-Z]+s?$" > words
```
