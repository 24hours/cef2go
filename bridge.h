#include <stdio.h>
#include "uthash.h" 

int lock = 0;

struct MemoryManagedBridge {
  void* key;        
  int count;
  void (*Deconstructor)(void *self);
  //Name          string
  UT_hash_handle hh;
};

struct MemoryManagedBridge *hash = NULL;

void create_handler(void* key) {
  struct MemoryManagedBridge *s;

  s = malloc(sizeof(struct MemoryManagedBridge));
  s->key = key;
  s->count = 0;
  HASH_ADD_PTR( hash, key, s );
}

struct MemoryManagedBridge* find_handler(void* key){
  struct MemoryManagedBridge *s = malloc(sizeof(struct MemoryManagedBridge));
  HASH_FIND_PTR(hash, key, s);  
  return s;
}

void delete_handler(void* key){
  struct MemoryManagedBridge* h = find_handler(key);
  HASH_DEL(hash, h);
  free(h);
}

void replace_handler(void* key, struct MemoryManagedBridge* rep){
  struct MemoryManagedBridge *s = malloc(sizeof(struct MemoryManagedBridge));
  HASH_REPLACE_PTR( hash, key, rep, s );
}

void acquire_lock(){
  while ( lock != 0 );
  lock = 1;
}

void release_lock(){
  lock = 0;
}

