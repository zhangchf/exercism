/**
 * Created by zhangchf on 28/08/2017.
 */
package Acronym

fun generate(phrase: String): String {
    if (phrase.contains(':')) {
        return phrase.substringBefore(':')
    }
    val splits = phrase.split(Regex("[ ,-]"))
    var acronym = StringBuilder()
    for(word in splits) {
        acronym.append(word.capitalize().filter { it in 'A'..'Z' })
    }
    return acronym.toString()
}