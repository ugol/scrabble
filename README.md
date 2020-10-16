# Scrabble
Words game helper, especially useful for Scrabble and similar games 

Basic usage
```
./words [letters] | sort -u
```

You can add a filter to get only words greater than something

```
./words [letters] 5 | sort -u
```

You can use capital letters: this means that those letters must be in the given position.
In this example you'll only get words with 'l' as first letter, 't' as third, and another 't' as fourth 

```
./words [LeTTers] 5 | sort -u
```

Default dictionary is in Italian

## Cloning and Building

```
git clone github.com/ugol/scrabble
cd scrabble
go get github.com/ugol/dwag

go build words.go
go build create-dictionary.go
```

# Example usage

Look for all words greater than 5 letters with e, s, e, m, p, i, o

```
> ./words esempio 5 | sort -u
emise
empie
empio
esempi
esempio
esime
esimo
esipo
espio
impose
ipomee
miope
omise
opime
poemi
poesi
poesie
poise
sepie
siepe
siepo
speme
spemi
```

Same example, with additional filtering for words starting with e and i as third letter
```
> ./words EsImpeo 5 | sort -u
emise
esime
esimo
esipo
```

# Improvements

Need to add support for jolly characters
Need to parametrize the dictionary file

# Creating the dictionary

./create-dictionary full-ita.txt full-ita.bin
