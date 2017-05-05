package primitive

/*
 * In most of the programming languages, collection of character yields a string, but go behaves differently
 *
 * https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/
 *
 * https://blog.golang.org/strings
 *
 * Strings are built from bytes so indexing them yields bytes, not characters.
 * A string might not even hold characters.
 * In fact, the definition of "character" is ambiguous and it would be a mistake to try to resolve the ambiguity by defining that strings are made of characters
 *
 */

const normalStr = "abc"

const specialStr = "abc本"
