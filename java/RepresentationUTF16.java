public class RepresentationUTF16 {
    public static void main(String[] args) {
        System.out.println("✋Hey".indexOf("H")); // 1
        System.out.println("🤚Hey".indexOf("H")); // 2
        System.out.println("✋Hey".charAt(1)); // H
        System.out.println("🤚Hey".charAt(1)); // ?
        System.out.println("✋Hey".codePointAt(1)); // U+0048 Latin Capital Letter H
        System.out.println("🤚Hey".codePointAt(1)); // U+DD1A Low Surrogate Code
    }
}
