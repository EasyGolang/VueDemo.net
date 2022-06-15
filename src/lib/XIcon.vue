<script setup lang="ts">
import { Icon } from '@vicons/utils';
import { defineAsyncComponent } from 'vue';

const props = defineProps({
  name: String,
});

function MyIconLoad() {
  switch (props.name) {
    case 'AudioFilled':
      return defineAsyncComponent(() => import('@vicons/antd/AudioFilled'));

    default:
      return defineAsyncComponent(() => import('@vicons/antd/CoffeeOutlined'));
  }
}

const MyIcon = MyIconLoad();
</script>

<template>
  <Icon>
    <MyIcon />
  </Icon>
</template>

<style lang="less"></style>
