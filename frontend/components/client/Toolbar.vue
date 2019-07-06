<template>
	<div class="
		bg-primary text-white
		rounded-b-card"
		style="background: linear-gradient(30deg, #FF6900, #FF974D)">
		<div class="p-4">
			<h1>Training title</h1>
		</div>
		<div class="p-4"
			title="Major graph">
			<trend 
				:data="major"
				:gradient="gradient"
				auto-draw
				smooth>
			</trend>
		</div>
	</div>
</template>

<script>
import Button from '../common/Button.vue';
import trend from 'vuetrend';

const gradient = ['#ffffff', '#ff974d'];

export default {
	components: {
		Button,
		trend,
	},
	data() {
		return {
		gradient,
		major: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0],
		}
	}
};
</script>
