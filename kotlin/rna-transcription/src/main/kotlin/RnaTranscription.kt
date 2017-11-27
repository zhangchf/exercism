fun transcribeToRna(dna: String): String = dna.map(::mapNucleocide).joinToString(separator = "")

val nucleotideMapping = mapOf('G' to 'C', 'C' to 'G', 'T' to 'A', 'A' to 'U')

fun mapNucleocide(c: Char): Char = nucleotideMapping.getOrElse(c) {throw RuntimeException("Invalid DNA Strand")}
