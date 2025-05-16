#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <cjson/cJSON.h>

typedef struct{
	char name[16];
	int amount;
	int *history;
	int history_size;
}temp_data;

temp_data create_temp(const char *name, int amount, int *history, int history_size){
	temp_data td;
	strncpy(td.name, name, sizeof(td.name) - 1);
	td.name[sizeof(td.name) - 1] = '\0';
	td.amount = amount;

	td.history = malloc(history_size * sizeof(int));
	if(td.history == NULL){
		perror("Failed to alloctae memory for history");
		exit(EXIT_FAILURE);
	}

	for (int i = 0; i < history_size; i++){
		td.history[i] = history[i];
	}

	td.history_size = history_size;
	return td;

}

void free_temp(temp_data *data, int *count){
	for(int i = 0; i <= *count; i++){
		free(data[i].history);
	}
	free(data);
}

temp_data* accessJSON(int *count){

	int size = 3;

	int history1[] = {10, 20, 30};
    	int history2[] = {5, 15, 25, 35};
    	int history3[] = {100, 200};
    
    	temp_data *list = malloc(size * sizeof(temp_data));
	if (list == NULL){
		*count = 0;
		return NULL;
	}
  	
	list[0] = create_temp("item1", 50, history1, sizeof(history1)/sizeof(history1[0]));
	list[1] = create_temp("item2", 75, history2, sizeof(history2)/sizeof(history2[0]));
	list[2] = create_temp("item3", 100, history3, sizeof(history3)/sizeof(history3[0]));

	*count = size;

	return list;
}

void showHistory(temp_data *list, int *count){
	
	int list_size = *count;

	for (int i = 0; i < list_size; i++){
		printf("Name: %s\n",list[i].name);
		printf("Amount: %d\n", list[i].amount);
		printf("History: ");
		for (int j = 0; j < list[i].history_size; j++){
			printf("%d ",list[i].history[j]);
		}
		printf("\n\n");
	}
}

void debugger(int *debug_count){
	printf("debug flag [%d]\n",*debug_count);
	*debug_count+=1;
}

int main(){
	int debug = 0;
	bool exit = false;
	debugger(&debug);//0

	int count;
	temp_data *data = accessJSON(&count);
	debugger(&debug);//1

	char buffer[100];
	int option;

	while(!exit){
	buffer[0] = '\0';
	option = -1;

	if(fgets(buffer,sizeof(buffer),stdin) == NULL){
		printf("Error reading input!\n");
		return 1;
	}
	if(sscanf(buffer, "%d", &option) == 0){
		printf("invalid input\n");
	}

	switch(option){
		case 0:
			exit = true;
			break;
		case 4:
			showHistory(data, &count);
			break;
		default:
			printf("out of scope number!\n");
			exit = true;
	}
	}

	free_temp(data,&count);
	
	debugger(&debug);//2

	return 0;
}
