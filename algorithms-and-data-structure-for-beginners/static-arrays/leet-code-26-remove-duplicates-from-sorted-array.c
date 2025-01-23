int removeDuplicates(int* nums, int numsSize) {
  const int OFFSET = 100; // porque os numeros vão de -100 a 100, não posso deixar índices negativos
  const int TAMANHO_INTERVALO = 201; // de -100 a 100 temos 201 números possíveis

  int k = 0;
  int numsQueJaApareceram[TAMANHO_INTERVALO] = {};

  for (int i = 0; i < numsSize; i++) {
    if (numsQueJaApareceram[nums[i]+OFFSET] != 1) {
      nums[k] = nums[i];
      k++;
      numsQueJaApareceram[nums[i]+OFFSET] = 1;
    }
  }

  return k;
}