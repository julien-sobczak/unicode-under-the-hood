import java.text.Normalizer;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class StringTest {

    private static void checkLength() {
        // Ʃ (https://unicode-table.com/en/01A9/)
        // This single character is encoded using at least two bytes
	    int l = "Ʃ".length();
	    if (l == 1) {
            System.out.println("👍 Unicode character length");
	    } else {
		    System.out.println("👎 Unicode character length. Got: " + l);
		}

        // Test with combining characters
        String sequence1 = "Voil\u00E0";
        String sequence2 = "Voil\u0061\u0300";
        if (sequence1.length() == sequence2.length()) {
            System.out.println("👍 Unicode sequence length");
        } else {
            System.out.println("👎 Unicode sequence length");
        }

        // String.length() is specified as returning the number of char values
        // ("code units") in the String.

        String norm1 = Normalizer.normalize(sequence1, Normalizer.Form.NFC);
        String norm2 = Normalizer.normalize(sequence2, Normalizer.Form.NFC);

        // Using code units
        if (norm1.length() == norm2.length()) {
            System.out.println("👍 Unicode sequence length using normalization and length()");
        } else {
            System.out.println("👎 Unicode sequence length using normalization and length()");
        }

        // Using code points
        int count1 = norm1.codePointCount(0, norm1.length());
        int count2 = norm2.codePointCount(0, norm2.length());
        if (count1 == count2) {
            System.out.println("👍 Unicode sequence length using normalization and codePointCount()");
        } else {
            System.out.println("👎 Unicode sequence length using normalization and codePointCount()");
        }
    }

    private static void checkComparison() {
        // Test with combining characters
        String sequence1 = "Voil\u00E0";
        String sequence2 = "Voil\u0061\u0300";
        if (sequence1.equals(sequence2)) {
            System.out.println("👍 Unicode string comparison using equals()");
        } else {
            System.out.println("👎 Unicode string comparison using equals()");
        }

        // Same problem as above. We need to normalize the string first.

        String norm1 = Normalizer.normalize(sequence1, Normalizer.Form.NFC);
        String norm2 = Normalizer.normalize(sequence2, Normalizer.Form.NFC);

        if (norm1.equals(norm2)) {
            System.out.println("👍 Unicode string comparison using normalization");
        } else {
            System.out.println("👎 Unicode string comparison using normalization");
        }
    }

    private static void checkCaseFolding() {
        String sequence1 = "Voil\u00E0"; // à
        String sequence2 = "Voil\u00C0"; // À

        String res = sequence1.toUpperCase();
        if (!res.equals(sequence2)) {
            System.out.println("👍 Unicode case folding using toUpperCase()");
        } else {
            System.out.println("👎 Unicode case folding using toUpperCase()");
        }

        res = sequence2.toLowerCase();
        if (!res.equals(sequence1)) {
            System.out.println("👍 Unicode case folding using toLowerCase()");
        } else {
            System.out.println("👎 Unicode case folding using toLowerCase()");
        }

        // No direct method to test Unicode case folding algorithm in Java
        // We need to use regexes instead. (next function)
    }

    private static void checkRegex() {
        List<String> res = new ArrayList<String>();
        String s = "100 µAh 10 mAh";
        Pattern p = Pattern.compile("\\d+ \\wAh");
        Matcher m = p.matcher(s);
        while (m.find()) {
          res.add(m.group());
        }
        if (res.size() == 2) {
            System.out.println("👍 Unicode regex using \\w");
	    } else {
		    System.out.println("👎 Unicode regex using \\w");
	    }

        // Like in Go, we need to use a special character class

        res = new ArrayList<String>();
        p = Pattern.compile("\\d+ \\p{L}Ah");
        m = p.matcher(s);
        while (m.find()) {
          res.add(m.group());
        }
        if (res.size() == 2) {
            System.out.println("👍 Unicode regex using \\p{L}");
	    } else {
		    System.out.println("👎 Unicode regex using \\p{L}");
	    }
    }

    public static void main(String[] args) {
        checkLength();
        checkComparison();
        checkCaseFolding();
        checkRegex();
    }
}

// Output:
// ------
// 1. Length
// 👍 Unicode character length
// 👎 Unicode sequence length
// 👍 Unicode sequence length using normalization and length()
// 👍 Unicode sequence length using normalization and codePointCount()
// 2. Text Comparison
// 👎 Unicode string comparison using equals()
// 👍 Unicode string comparison using normalization
// 3. Text Case Folding
// 👍 Unicode case folding using toUpperCase()
// 👍 Unicode case folding using toLowerCase()
// 4. Regex
// 👎 Unicode regex using \w
// 👍 Unicode regex using \p{L}
