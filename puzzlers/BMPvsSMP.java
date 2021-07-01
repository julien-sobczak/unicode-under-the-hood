public class BPMvsSMP {
    public static void main(String[] args) {
        System.out.println("✋Hey".indexOf("H")); // 1
        System.out.println("🤚Hey".indexOf("H")); // 2
        System.out.println("✋Hey".charAt(1)); // H
        System.out.println("🤚Hey".charAt(1)); // ?

        // Why
        // ✋ RAISED HAND U+270B is in the BMP (2 bytes in UTF-16 => size 1)
        // 🤚 RAISED BACK OF HAND U+1F91A is in the SMP (and thus need a surrogate pair in UTF-16 => size 2)
    }
}
