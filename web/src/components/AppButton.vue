<template>
  <div class="btn">
    {{ text }}
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "AppButton",
  props: {
    text: {
      type: String,
      required: true,
    },
  },
});
</script>

<style scoped>
.btn {
  cursor: pointer;
  border: none;
}
</style>
