# Building an xxd clone

#### Why?
We work with binary files all the time, sometimes not even realising we when we are. I've been wanting to dive in and get a deeper understanding of binary files for a while, and John's challenge seemed to pop up at a good time. I've got two main goals during this challenge. First is to understand and recognise the output of `xxd`. Offset, hex data, and text data. Secondly, attempt to be able to do data manipulation by modifying hex dump. This is just a nice to have though.
#### The Challenge
I've broken this challenge down similarly to how John has laid out the steps.
- Setup the project, read user input and output hex somewhat close to how xxd does.
- Support little endian and grouping flag.
- Support number of octets.
- Implement seek.
- Revert hex back to binary.

#### Step 1
Setting up the project usually goes the same way for me, especially for these challenges. I'll make sure I can read args and handle files.

For step one, we only need to read in the arg for a file path. We bottom out if we don't get an arg.

Next we'll open the file passed to us and add our defer.

```Go
if len(os.Args) < 1 {
	os.Exit(1)
}

file, err := os.Open(os.Args[1])
if err != nil {
	log.Fatal(err)
}
defer file.Close(
```

Now we'll work on the main part of what's needed for step one. Reading the input and dumping the hex.

I tried out a few different approaches here, specifically around the program loop. I'm not a huge fan of unconditional loops but here I think it works. What we're doing is breaking out of the loop once we hit an EOF error. I guess it's not totally unconditional but it's not too elegant. Something we can return to later however.

Once we know we're good, we start by creating a buffer for `fileBytes`. From here on out is pretty new territory for me. I've not really worked with `binary.Read`, or even `hex.Dump` before.

Looking at `binary.Read`, we pass in our buffer from the line above. Followed by the byte order, there is a step ahead which requires us to handle for little endian but we'll leave it as big endian for now. Finally, we'll pass our slice where we want our output to live.

```Go
buffer := bytes.NewBuffer(fileBytes)
err = binary.Read(buffer, binary.BigEndian, &out)
```

For the output, we can use `hex.Dump` which requires a byte slice. The output data is correct, or at least matches what `xxd` spits out. The formatting is different, but again, we can come back to this.

```Go
fmt.Printf("%s", hex.Dump(out))
```

*Output step 1 image here*

---
