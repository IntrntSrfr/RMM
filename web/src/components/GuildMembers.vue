<template>
  <div class="options">
    <AppButton text="Select all" @click="markAll" />
    <AppButton text="Select none" @click="markNone" />
    <AppButton text="Ban selected" @click="banSelected" />
  </div>
  <table>
    <thead>
      <tr>
        <th>
          <input
            type="checkbox"
            :checked="selectedMembers.length===members.length"
            :indeterminate="
              selectedMembers.length &&
              selectedMembers.length !== members.length
            "
            @click="selectedMembers.length ? markAll() : markNone()"
          />
        </th>
        <th>#</th>
        <th>ID</th>
        <th>Username</th>
        <th>Joined</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(m, i) in members" :key="i">
        <td>
          <input type="checkbox" v-model="m.checked" />
        </td>
        <td>{{ i }}</td>
        <td>{{ m.member.user.id }}</td>
        <td>
          {{ `${m.member.user.username}#${m.member.user.discriminator}` }}
        </td>
        <td>5h ago</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { Member } from "@/stores/user";
import { useUserStore } from "@/stores/user";
import axios from "axios";
import AppButton from "@/components/AppButton.vue";

type TableMember = {
  member: Member;
  checked: boolean;
};

export default defineComponent({
  name: "GuildMembers",
  components: { AppButton },
  setup() {
    const userStore = useUserStore();
    return { userStore };
  },
  data() {
    return {
      members: [] as TableMember[],
    };
  },
  computed: {
    selectedMembers(): TableMember[] {
      return this.members.filter((m: TableMember) => m.checked);
    },
  },
  async created() {

    const guildID = this.$route.params.guildID;
    if (!guildID) {
      return;
    }

    try {
      const res = await axios.get<Member[]>(
        "http://localhost:4444/api/guilds/" + guildID + "/members"
      );
      this.members = res.data
        .reverse() // we want the latest entries at the top
        .map((m) => ({ checked: false, member: m } as TableMember));
    } catch (error) {
      console.log(error);
    }
  },

  methods: {
    markAll() {
      this.members.forEach((m: TableMember) => (m.checked = true));
    },
    markNone(): void {
      this.members.forEach((m: TableMember) => (m.checked = false));
    },
    banSelected() {
      const ans = confirm("Are you SURE you want to ban ALL selected users?");
      if (!ans) {
        return;
      }

      console.log(this.selectedMembers);
    },
  },
});
</script>

<style scoped>
.options {
  display: flex;
  gap: 1em;
}
table {
  width: 100%;
  border: 1px solid dodgerblue;
}
thead {
  text-align: left;
  border-bottom: 1px solid dodgerblue;
}
tr:nth-child(2n) {
  background-color: #282828;
}
</style>
