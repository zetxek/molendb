# So... what is this? #

This project lets you check the closest windmill to a given point, in The Netherlands.

![vincent-versluis-198649.jpg](https://bitbucket.org/repo/LooGk8B/images/1161534428-vincent-versluis-198649.jpg)



## Why?!

Just because. And because I was moving from one apartment to Amsterdam to another, and I was curious about which place would have a windmill closer (https://twitter.com/zetxek/status/865320654786220037). It turned to be the 1st apartment!

```
go run main.go
637 De Otter 0.42759532
635 De Bloem 0.75619304
```

## How?

Thanks to the great informatoin from https://molendatabase.nl.
The project downloads the files (to not overload the servers or cause any extra load on their site) and creates a local db with the windmills.
Then, the distance calculation is done with a SQLite query.

## What else?

Well, as the idea was playing a little bit with golang, let's organize the repo better.