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

Use `-hints` to only get the start of the answers.

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

or when using `-hints` only:

```
20:48:09.648 [INF] Words in dictionary : 132446
20:48:09.648 [INF] Looking at fragments:
lb	mac	ash	pa
tri	ns	ar	arch
ect	fl	el	oo
bu	ora	al	tes
pre	dic	tab	le
20:48:09.648 [INF] Checking 2 parts combinations:
al-*	: al__
bu-*	: bu__
bu-*	: bu__
el-*	: el___
fl-*	: fl___
fl-*	: fl___
le-*	: le__
le-*	: le___
le-*	: le__
mac-*	: mac__
ora-*	: ora___
pa-*	: pa__
pa-*	: pa__
pa-*	: pa___
tab-*	: tab__
tab-*	: tab__
tri-*	: tri__
20:48:09.648 [INF] Checking 3 parts combinations:
bu-*-*	: bu____
le-*-*	: le____
pa-*-*	: pa_______
pre-*-*	: pre_____
tes-*-*	: tes_____
tri-*-*	: tri_____
20:48:09.651 [INF] Checking 4 parts combinations:
el-*-*-*	: el_________
fl-*-*-*	: fl_______
mac-*-*-*	: mac______
pa-*-*-*	: pa_________
pre-*-*-*	: pre________
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
  -hints
    	Only show the first segment of the answers (hints only mode)
```

## What is the Quartiles game?

https://www.apple.com/newsroom/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/


<img src="https://github.com/ldemailly/quartiles/assets/3664595/4ef37a1b-86e8-4841-88fa-c36bc5e5838b" width="50%">

## Wordlist

The `words` used in Dockerfile is from Ubuntu's
`apt install wamerican-large` american-english-large words-large
removing 's and all caps acronyms and non ascii words
```
gnugrep -v -E -e "^[A-Z]+s?$" -e '[^a-zA-Z]' < /usr/share/dict/american-english-large > words
```
