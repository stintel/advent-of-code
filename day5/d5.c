#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int get_col(char* s) {
	uint8_t col = 0;

	for (uint8_t i = 0; i < 3; i++) {
		if (s[i] == 'R')
			col |= 1 << (2 - i);
	}

	return(col);
}

int get_row(char* s) {
	uint8_t row = 0;

	for (uint8_t i = 0; i < 7; i++) {
		if (s[i] == 'B')
			row |= 1 << (6 - i);
	}

	return(row);
}

int get_id(int r, int c) {
	return (r * 8 + c);
}

int main() {
	FILE *fp;
	char *line = NULL;
	char str_col[4], str_row[9];
	size_t len = 0;
	ssize_t read;
	uint8_t col, row = 0;
	uint16_t id, maxid = 0;
	uint16_t seats[128][8];

	fp = fopen("input", "r");
	if (fp == NULL)
		exit(EXIT_FAILURE);

	while ((read = getline(&line, &len, fp)) != -1) {
		strncpy(str_row, line, 7);
		strncpy(str_col, line + 7, 3);

		row = get_row(str_row);
		col = get_col(str_col);
		id = get_id(row, col);

		seats[row][col] = 1;

		if (id > maxid)
			maxid = id;
	}

	fclose(fp);
	free(line);

	printf("Highest seat id: %i\n", maxid);

	for (int i = 0; i < 128; i++) {
		for (int j = 0; j < 8; j++) {
			if (seats[i][j] != 1) {
				id = get_id(i, j);

				if (seats[i][j-1] == 1 && seats[i][j+1] == 1)
					printf("Your seat id is %i\n", id);
			}
		}
	}
}
