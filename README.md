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
$ quartiles "nder
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
10:40:06.232 [INF] Log level is now 0 Debug (was 2 Info)
10:40:06.274 [INF] Words in dictionary /usr/share/dict/words: 234456
10:40:06.274 [INF] Looking at fragments:
nder	nt	sti	bi
lo	thu	al	con
ies	vo	mon	ba
tue	rd	fri	lit
sw	ck	wed	gam
10:40:06.274 [INF] Checking 2 words combinations:
alba
albi
almon
baal
babi
back
bacon
bander
bant
bard
bick
bilo
binder
bint
bird
bisti
conal
gamba
gammon
lock
lord
monal
stick
stint
thunder
vowed
10:40:06.274 [INF] Checking 3 words combinations:
stibial
swallo
wedlock
10:40:06.276 [INF] Checking 4 words combinations:
backgammon
birdbander
constituent
thunderbird
thunderstick
```

## Usage

```
quartiles 0.1.0 usage:
	quartiles [flags] word fragments...
Finds all words in dictionary that can be made from the given fragments
or 1 of the special arguments
	quartiles {help|envhelp|version|buildinfo}
flags:
  -dictionary path
    	Dictionary file path to use (default "/usr/share/dict/words")
```

## What is the Quartiles game?

https://www.apple.com/newsroom/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/

![screenshot](https://www.apple.com/newsroom/images/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/article/Apple-News-Plus-Quartiles-initial-state_inline.jpg.large_2x.jpg)
