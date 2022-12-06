import os


class CodeSequence:
    """
    Represents a code sequence.
    """

    def __init__(self):
        """
        Create a new CodeSequence instance.
        """
        self._code_sequence_start_of_packet = []
        self._code_sequence_start_of_message = []

    def update(self, symbol):
        """
        Update the sequence with the new symbol.
        :param symbol: The symbol to add.
        :return: Tuple of Booleans (a, b) - if a is True, we have a start of packet, if b is True, we have a start of
                 message.
        """
        self._code_sequence_start_of_packet.append(symbol)
        self._code_sequence_start_of_message.append(symbol)
        if len(self._code_sequence_start_of_packet) > 4:
            self._code_sequence_start_of_packet.pop(0)
        if len(self._code_sequence_start_of_message) > 14:
            self._code_sequence_start_of_message.pop(0)
        test_set_start_of_packet = set(self._code_sequence_start_of_packet)
        test_set_start_of_message = set(self._code_sequence_start_of_message)
        return len(test_set_start_of_packet) == 4, len(test_set_start_of_message) == 14

    def __str__(self):
        """
        Get the current code sequence as a string.
        :return: The current code sequence as a string.
        """
        sequence_start_of_packet = ''
        sequence_start_of_message = ''
        for symbol in self._code_sequence_start_of_packet:
            sequence_start_of_packet = '{}{}'.format(sequence_start_of_packet, symbol)
        for symbol in self._code_sequence_start_of_message:
            sequence_start_of_message = '{}{}'.format(sequence_start_of_message, symbol)
        return '{}/{}'.format(sequence_start_of_packet, sequence_start_of_message)


def day06_main(input_path=None):
    """
    Solve day 06 of Advent of Code 2022.
    :param input_path: Path to the input data. If None, attempt to autodetect the path from this Python file name.
    :return: None.
    """
    if input_path is None:
        own_file_name = os.path.splitext(os.path.basename(__file__))[0]
        input_path = '../../data/{}/input'.format(own_file_name)
    with open(input_path) as input_file:
        code_line = input_file.readline()
        code_sequence = CodeSequence()
        symbol_count = 0
        start_of_packet_found = False
        start_of_message_found = False
        for symbol in code_line:
            symbol_count += 1
            start_of_packet, start_of_message = code_sequence.update(symbol)
            if start_of_packet and not start_of_packet_found:
                print("Start-of-Packet sequence found at index {}: {}".format(symbol_count, code_sequence))
                start_of_packet_found = True
            if start_of_message and not start_of_message_found:
                print("Start-of-Message sequence found at index {}: {}".format(symbol_count, code_sequence))
                start_of_message_found = True
            if start_of_packet_found and start_of_message_found:
                break


if __name__ == '__main__':
    day06_main()
