# __Zypher__ 

# Description
This is a Go library I built. It is an encryption tool that has *4* different methods for generating hashes for plaintext. I wanted to create this library so I had tools I could use to help lower the chances of hash collisions and makes the it harder to determine the hashing algorithm used. At a highlevel the methods take in plaintext and cyphers the characters in it then hashes the result. There are a few other varaitions you can add like number of times to run the text through a cypher, number of times to hash the result, alternate adding and subtracting the shift, and ignoring spaces. I got the idea for repeatedly cyphering and hashing the plaintext from studying blockchain and learning how the pub/priv keys are generated. The four methods available in the library are: 
- *AsciZyph* : Cyphers the plaintext x number of times then hashes the result n number of times. Only works on alpha numeric strings
- *HexZyph* : Cyphers the plaintext x number of times then hashes the result n number of times. Only works on hex strings
- *Zyph* : Cyphers the plaintext x number of times then hashes the result n number of times. Works with alpha numeric strings containing symbols no emojis
- *ZypHash* : Hashes plain text x number of times. Works with alpha numeric strings containing symbols

*Each method call uses values from the configured Zypher calling the method. I.E. creating a new Zypher with NewZypher() creates a default Zypher instance, calling Zyph will cypher the plaintext 3 times then hash the result 3 times. Any method that cyphers x number of times or hashes n number of times will use the values defined in the Zypher instance*


# Quickstart
There are two ways to use the methods.
- *Option 1* - Import the library from my repository, call the NewZypher method that creates a default Zypher object then pass plaintext into anyone of the 4 methods.

- *Option 2* - To pull down the code to be able to run the test or run locally. Clone the repository, then run the command `go test` or create a main.go file, use the methods the run the command `go run .`

# Contribute
To contribute pulldown the repo, after making any changes create a pull request.
