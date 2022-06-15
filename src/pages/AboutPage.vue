<script setup lang="ts">
import { getToken, mStorage } from '@/utils/tools';
import { outLogin } from '@/api/Account';
import XIcon from '@/lib/XIcon.vue';

const token = getToken();
const UserInfo = mStorage.get('user_info');

const handleOutLogin = () => {
  outLogin();
};
</script>

<template>
  <div class="about">
    <XIcon color="green" size="48" name="AudioFilled" />

    <div v-if="token">
      <br />
      <n-avatar
        round
        :size="100"
        :src="UserInfo.Avatar"
        fallback-src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
      />
      <br />
      <div>欢迎访问,{{ UserInfo.NickName }}</div>
      <br />
      <n-button @click="handleOutLogin">退出登录</n-button>
    </div>
    <div v-if="!token">
      <br />
      <a href="/login"> <n-button>登录</n-button> </a>
      <br />
    </div>
  </div>
</template>

<style lang="less">
.Avatar {
  width: 100px;
  height: 100px;
  border-radius: 100px;
}
</style>
