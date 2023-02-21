<template>
  <div class="member-wrapper">
    <p class="counter">
      {{ selectedMembers.length }} / {{ members.length }} selected
    </p>
    <div class="inp-text">
      <input type="text" v-model="search" />
      <div
        class="inp-text__option"
        title="Use regular expressions"
        @click="() => (useRegex = !useRegex)"
        :class="{ checked: useRegex }"
      >
        .*
      </div>
    </div>
    <div class="actions">
      <div class="grp">
        <AppButton text="Select none" @click="markNone" />
        <AppButton text="Select all" @click="markAll" />
      </div>
      <div class="grp">
        <AppButton text="Ban selected" variant="danger" @click="banSelected" />
      </div>
    </div>
    <AppLoader v-if="isLoadingMembers" />
    <div v-else class="member-list">
      <GuildMembersItem
        v-for="(m, i) in searchedMembers"
        :key="i"
        @click="mark(m)"
        :checked="m.checked"
        :id="m.member.user.id"
        :username="`${m.member.user.username}#${m.member.user.discriminator}`"
        :icon="m.member.user.avatar"
        :joined="m.joined"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import type { Member } from "@/stores/user";
import { useGuildStore } from "@/stores/guilds";
import AppButton from "@/components/AppButton.vue";
import GuildMembersItem from "./GuildMembersItem.vue";
import AppLoader from "./AppLoader.vue";

type TableMember = {
  member: Member;
  checked: boolean;
  joined: Date;
};

const guildStore = useGuildStore();
const route = useRoute();
const search = ref("");
const useRegex = ref(false);
const members = ref<TableMember[]>([]);

const selectedMembers = computed(() => {
  return members.value.filter((m: TableMember) => m.checked);
});

const searchedMembers = computed(() => {
  if (!search.value) return members.value;
  return members.value.filter((m: TableMember) =>
    useRegex.value
      ? m.member.user.username.match(search.value)
      : m.member.user.username
          .toUpperCase()
          .includes(search.value.toUpperCase())
  );
});

const isLoadingMembers = computed(() => {
  const guildID = route.params.guildID;
  const g = guildStore.getGuildByID(guildID.toString());
  if (!g) return true;
  return !g.members;
});

const markAll = () => {
  searchedMembers.value.forEach((m) => (m.checked = true));
};

const markNone = () => {
  searchedMembers.value.forEach((m) => (m.checked = false));
};

const mark = (tm: TableMember) => {
  tm.checked = !tm.checked;
};

const banSelected = () => {
  if (!selectedMembers.value.length) return;

  const ans = confirm("Are you SURE you want to ban ALL selected users?");
  if (!ans) return;

  console.log(selectedMembers.value.map((m) => m.member.user.id).join(" "));
};

const OnCreated = async () => {
  const guildID = route.params.guildID;
  const g = guildStore.getGuildByID(guildID.toString());
  await guildStore.fetchGuildMembers(guildID.toString());

  if (!g || !g.members) return;

  members.value = g.members
    .map((m: Member) => {
      return {
        member: m,
        checked: false,
        joined: new Date(m.joined_at),
      } as TableMember;
    })
    .sort((a: TableMember, b: TableMember) => {
      return b.joined.getTime() - a.joined.getTime();
    });
};

OnCreated();
</script>

<style scoped>
.member-wrapper {
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.actions {
  display: flex;
  justify-content: space-between;
  gap: 1em;
  padding-bottom: 1em;
  border-bottom: 1px solid #666;
}

.grp {
  display: flex;
  gap: 1em;
}

.member-list {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.member + .member {
  border-top: 1px solid dodgerblue;
}

.inp-text {
  display: flex;
  align-items: center;
  padding: 0.25em;

  background-color: #191b1f;
  border-radius: 0.25em;
}
.inp-text input {
  flex-grow: 1;
  background: transparent;
  appearance: none;
  border: none;
  outline: none;

  font-size: 1rem;
  color: rgb(172, 172, 172);
}
.inp-text__option {
  cursor: pointer;
  display: flex;
  justify-content: center;
  width: 24px;
  height: 24px;
  border-radius: 0.25em;
}
.inp-text__option:hover {
  background-color: #26292f;
}
.inp-text__option.checked {
  background-color: #32373d;
}
</style>
