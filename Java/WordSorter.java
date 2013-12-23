public class WordSorter extends Thread implements WordHandler {

    private String[] sorted = null;
    private String[] sourceStrings = null;
    private int start = 0;
    private int end = 0;


    public WordSorter(String[] sourceList, int start, int end) {
        this.sourceStrings = sourceList;
        this.start = start;
        this.end = end;

        start();
    }


    @Override
    public void run() {
        SortedStringsList sorter = new SortedStringsList();

        for (int i = start; i < end; i++) {
            sorter.add(sourceStrings[i]);
        }

        sorted = sorter.toArray();
    }


    public String[] getWords() {
        return sorted;
    }
}

