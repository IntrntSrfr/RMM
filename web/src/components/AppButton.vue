<template>
  <div class="btn" :class="{ outlined: variant === 'outlined' }">
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
    variant: {
      type: String,
      default: "outlined",
      validator(value: string) {
        return ["outlined", "filled", "danger"].includes(value);
      },
    },
  },
});
</script>

<style scoped>
.btn {
  position: relative;
  cursor: pointer;
  border: none;
  font-size: 1rem;
  font-family: Inter, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    Oxygen, Ubuntu, Cantarell, "Fira Sans", "Droid Sans", "Helvetica Neue",
    sans-serif;
  padding: 0.25em 0.5em;
  border-radius: 5px;
  transition: 0.2s;
}

.outlined {
  --text-color: white;
  --border-color: dodgerblue;
  --border-hover-color: rgba(22, 115, 213, 0.2);
  color: var(--text-color);
}

.outlined:after {
  content: "";
  position: absolute;
  height: 100%;
  width: 100%;
  left: 0;
  top: 0;
  border: 1px solid var(--border-color);
  border-radius: 5px;
}

.outlined:hover {
  background-color: var(--border-hover-color);
}
</style>
