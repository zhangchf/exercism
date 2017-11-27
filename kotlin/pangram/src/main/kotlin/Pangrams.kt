/**
 * Created by zhangchf on 25/08/2017.
 */
package Pangrams

/*
fun isPangram(sentence: String): Boolean =
    sentence.fold(HashMap<Char, Int>(), { map, char ->
        if (char.toLowerCase() in 'a'..'z') {
            map[char] = 1
        }
        map
    }).values.size == 26
*/

fun isPangram(sentence: String): Boolean =
        sentence
                .asSequence()
                .filter { it.isLetter() }
                .map { it.toLowerCase() }
                .distinct()
                .count() == 26

