<template>
<section class="wf-canvas">
	<div class="wf-canvas-inner">
		<div class="wf-canvas-body dropbody" 
			v-bind:class="{empty:isEmpty}">
			<div class="wf-dragging-mark" v-if="inCanvas==0"></div>
			<div v-for="(item,index) in components" :key="index">
				<!-- textarea -->
				<div v-if="item.componentName=='textareafield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component wf-component-textarea">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- text-->
				<div v-if="item.componentName=='textfield'">
					<div 
						@click="current"
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- number -->
				<div v-if="item.componentName=='numberfield'">
					<div 
						@click="current"
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- select -->
				<div v-if="item.componentName=='selectfield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-arrow-right"></i>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- multi select -->
				<div v-if="item.componentName=='multiselectfield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-arrow-right"></i>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- date -->
				<div v-if="item.componentName=='datefield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-arrow-right"></i>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- date-range -->
				<div v-if="item.componentName=='daterangefield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component wf-component-daterange">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
									{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-arrow-right"></i>
								</span>
							</div>
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
									{{item.defaultLable2}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps2}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-arrow-right"></i>
								</span>
							</div>
							<div class="wf-componentview-border" 
								v-if="item.defaultAutorekonTime">
								<label class="wf-componentview-label">
									<span>{{item.defaultSubtitle}}</span>
									<span>（天）</span>
								</label>
								<span class="wf-componentview-placeholder"></span>
							</div>
							<span></span>
						</div>
					</div>
					<div class="wf-dragging-mark" v-if="inCanvas==index+1"></div>
				</div>
				<!-- picture -->
				<div v-if="item.componentName=='photofield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-camera-solid icon-camera"></i>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- note -->
				<div v-if="item.componentName=='textnote'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component wf-component-textnote">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- attachment -->
				<div v-if="item.componentName=='attachment'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
									<i class="el-icon-paperclip icon-paperclip"></i>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- identity-->
				<div v-if="item.componentName=='identityfield'">
					<div 
						@click="current"
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
								{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
									<span v-if="item.defaultRequired">（必填）</span>
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- phone-->
				<div v-if="item.componentName=='phonefield'">
					<div 
						@click="current"
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component wf-component-phone">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
									<b v-if="item.defaultRequired">*</b>
									{{item.defaultLable}}
								</label>
							</div>
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
									<span>+86</span>
									<i class="el-icon-arrow-right"></i>
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- location -->
				<div v-if="item.componentName=='locationfield'">
					<div 
						@click="current"
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
									{{item.defaultLable}}
								</label>
								<span class="wf-componentview-placeholder">
									{{item.defaultProps}}
								</span>
							</div>
							<div class="wf-componentview-tips">当前时间</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- rate -->
				<div v-if="item.componentName=='ratefield'">
					<div 
						@click="current"
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component wf-component-rate">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<div class="wf-componentview-border">
								<label class="wf-componentview-label">
									{{item.defaultLable}}
								</label>
							</div>
							<div class="wf-componentview-border">
								<span class="wf-componentview-rate"
									v-bind:class="item.defaultRate==5?'rate20':'rate10'"
									v-for="i in parseInt(item.defaultRate)"
									:key="i">
									<i class="el-icon-star-on"></i>
									{{i}}
								</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
				<!-- table -->
				<div v-if="item.componentName=='tablefield'">
					<div 
						@click="current" 
						draggable="true" 
						v-on:dragstart="moveStart"
						v-on:mouseout="mouseOut"
						v-on:mouseover="hover"
						:data-index="index"
						v-bind:class="index==selected?'active':''"
						class="wf-component wf-component-tablefield">
						<div class="wf-remove" 
							@click="close">
							<i class="el-icon-close"></i>
						</div>
						<div class="wf-overlay"></div>
						<div class="wf-componentview">
							<label class="wf-componentview-title">
								{{item.defaultLable}}
							</label>
							<div class="wf-componentview-area" 
								v-bind:class="{empty:item.components.length<=0}">
								<span class="emptytip" 
									v-if="item.components.length<=0">可拖入多个组件（不包含明细组件）
								</span>
								<div class="wf-component-group dropbody">
									<div class="wf-dragging-mark" 
										v-if="item.InTableCanvas==0"></div>
									<div v-for="(itemInTable,idx) in item.components"
										:key="idx">	
										<!-- table textarea -->
										<div 
											v-if="itemInTable.componentName=='textareafield'">
											<div 
												@click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component wf-component-textarea">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">		{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table text -->
										<div v-if="itemInTable.componentName=='textfield'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">		{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table number -->
										<div v-if="itemInTable.componentName=='numberfield'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table select -->
										<div v-if="itemInTable.componentName=='selectfield'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseover="hover"
												v-on:mouseout="mouseOut"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
														{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="el-icon-arrow-right"></i>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table multi select -->
										<div v-if="itemInTable.componentName=='multiselectfield'">
											<div 
												@click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
														{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">	 {{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="el-icon-arrow-right"></i>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table date -->
										<div v-if="itemInTable.componentName=='datefield'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="el-icon-arrow-right"></i>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table daterange -->
										<div 
											v-if="itemInTable.componentName=='daterangefield'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component wf-component-daterange">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="el-icon-arrow-right"></i>
														</span>
													</div>
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable2}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps2}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="el-icon-arrow-right"></i>
														</span>
													</div>
													<div class="wf-componentview-border"
														 v-if="itemInTable.defaultAutorekonTime">
														 <label class="wf-componentview-label">
															 	<span>
															 		{{itemInTable.defaultSubtitle}}
															 	</span>
														 		<span>（天）</span>
														 </label>
														 <span class="wf-componentview-placeholder"></span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table photo -->
										<div v-if="itemInTable.componentName=='photofield'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="icon-camera"></i>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table textnote -->
										<div v-if="itemInTable.componentName=='textnote'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component wf-component-textnote">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table attachment -->
										<div v-if="itemInTable.componentName=='attachment'">
											<div @click="current" 
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">（必填）
															</span>
															<i class="el-icon-paperclip icon-paperclip"></i>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table identity -->
										<div v-if="itemInTable.componentName=='identityfield'">
											<div 
												@click="current"
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
														{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
															<span v-if="itemInTable.defaultRequired">
															（必填）
															</span>
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table phone -->
										<div v-if="itemInTable.componentName=='phonefield'">
											<div 
												@click="current"
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component wf-component-phone">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															<b v-if="itemInTable.defaultRequired">*</b>
															{{itemInTable.defaultLable}}
														</label>
													</div>
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															<span>+86</span>
															<i class="el-icon-arrow-right"></i>
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table location -->
										<div v-if="itemInTable.componentName=='locationfield'">
											<div 
												@click="current"
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component">
												<div class="wf-remove" @click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
														{{itemInTable.defaultLable}}
														</label>
														<span class="wf-componentview-placeholder">
															{{itemInTable.defaultProps}}
														</span>
													</div>
													<div class="wf-componentview-tips">当前时间</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
										<!-- table rate -->
										<div v-if="itemInTable.componentName=='ratefield'">
											<div 
												@click="current"
												draggable="true" 
												v-on:dragstart="moveStart"
												v-on:mouseout="mouseOut"
												v-on:mouseover="hover"
												:data-index="idx"
												v-bind:class="idx==item.selected?'active':''"
												class="wf-component wf-component-rate">
												<div class="wf-remove" 
													@click="close">
													<i class="el-icon-close"></i>
												</div>
												<div class="wf-overlay"></div>
												<div class="wf-componentview">
													<div class="wf-componentview-border">
														<label class="wf-componentview-label">
															{{itemInTable.defaultLable}}
														</label>
													</div>
													<div class="wf-componentview-border">
														<span class="wf-componentview-rate"
															v-bind:class="itemInTable.defaultRate==5?'rate20':'rate10'"
															v-for="i in parseInt(itemInTable.defaultRate)"
															:key="i">
															<i class="el-icon-star-on"></i>
															{{i}}
														</span>
													</div>
												</div>
											</div>
											<div class="wf-dragging-mark" 
												v-if="item.InTableCanvas==idx+1"></div>
										</div>
									</div>
								</div>
							</div>
							<div class="wf-componentview-adddetail">
								<i class="el-icon-plus"></i>
								<span>{{item.defaultAction}}</span>
							</div>
						</div>
					</div>
					<div class="wf-dragging-mark" 
						v-if="inCanvas==index+1"></div>
				</div>
			</div>
		</div>
	</div>
</section>
</template>
<script>
export default{
  name: 'wf-canvas',
  props: {
    list: {
      type: Array,
      default: null
    }
  },
  data() {
		return {
			frame: null,
			frameScrollTop: 0,
			canvas: null,
			canvasScrollTop: 0,
			isEmpty: true,
			left: 0,
			top: 0,
			scrollTop: 0,
			width: 0,
			height: 0,
			inCanvas: null,
			components: [],
			selected: null,
			domArr: [],
			isDrag: false,     // 手机内组件是否被拖动
			dragIndex: null,   // 被拖动组件的index
			tabIndex: null,    // 拖动时table组件的index
			parNodeIndex: null // 被拖动组件的table index
		}
  },
  methods: {
		hover(e) {
			e.stopPropagation()
			e.preventDefault()
			e.currentTarget.classList.add('hover')
		},
		mouseOut(e) {
			e.stopPropagation()
			e.preventDefault()
			e.currentTarget.classList.remove('hover')
		},
		close(e) {
			e.preventDefault()
			e.stopPropagation()
			let dom = e.currentTarget.parentNode
			let parNode = dom.parentNode.parentNode.parentNode.parentNode
			let index = dom.getAttribute('data-index')
			// 如果是table
			if (parNode.className.indexOf('wf-componentview-area') >= 0) {
				let parNodeIndex = parNode.parentNode.parentNode.getAttribute('data-index')
				this.components[parNodeIndex].components.splice(index, 1)
				if (parNode.className.indexOf('active') > 0) {
					this.selected = null
					this.components[parNodeIndex].selected = null
				}
			} else {
				this.components.splice(index, 1)
				// 是否是active项  删除需要设置isSetting=false
				if (dom.className.indexOf('active') > 0) {
					this.selected = null
					drag.$emit("showSetting", false)
				} 
				if (this.components.length <= 0) {
					this.isEmpty = true
				}
			}
		},
		current(e) {
			e.stopPropagation()
			let dom = e.currentTarget
			let index = dom.getAttribute('data-index')
			let parNode = dom.parentNode.parentNode.parentNode.parentNode
			let obj
			if (parNode.className.indexOf('wf-componentview-area') >= 0) {
				let parNodeIndex = parNode.parentNode.parentNode.getAttribute('data-index')
				obj = this.components[parNodeIndex].components[index]
				if (this.selected == index && this.parNodeIndex != parNodeIndex || this.selected !== index) {
					this.parNodeIndex = parNodeIndex
					this.selected = null
					this.components[parNodeIndex].selected = index
					drag.$emit("selectComponent", obj)
				}
			} else {
				this.parNodeIndex = null
				for (let i = 0, l = this.components.length; i < l; i++) {
					this.components[i].selected = null
				}
				obj = this.components[index]
				if (this.selected !== index) {
					this.selected = index
					drag.$emit("selectComponent", obj)
				}
			}
		},
		moveStart(e) {
			e.preventDefault()
			e.stopPropagation()
			let obj = {}
			let dom = e.currentTarget
			dom.classList.add('draging')
			let index = dom.getAttribute('data-index')
			let actualLeft = dom.offsetLeft
			let actualTop = dom.offsetTop
			let current = dom.offsetParent
			while (current !== null) {
				actualLeft += current.offsetLeft
				actualTop += current.offsetTop
				current = current.offsetParent
			}
			obj.clientX = e.clientX
			obj.clientY = e.clientY
			obj.isStart = true
			let parNode = dom.parentNode.parentNode.parentNode.parentNode
			
			if (parNode.className.indexOf('wf-componentview-area') >= 0) {
				let parNodeIndex = parNode.parentNode.parentNode.getAttribute('data-index')
				obj.componentView = this.components[parNodeIndex].components[index]
				obj.componentName = this.components[parNodeIndex].components[index].componentName
				obj.componentText = this.components[parNodeIndex].components[index].name
				this.parNodeIndex = parNodeIndex 
			} else {
				obj.componentView = this.components[index]
				obj.componentName = this.components[index].componentName
				obj.componentText = this.components[index].name
				this.parNodeIndex = null
			}

			this.isDrag = true
			this.dragIndex = index

			drag.$emit("moveStart", obj)
		},
		addComponents(obj, list, sort) {
			let self = this
			let component = JSON.parse(JSON.stringify(obj))
			let componentsLength = self.getComponents()
			
			let idx = 0
			for (let i = 0; i < list.length; i++) {
				let item = list[i]
				if (item.name == component.componentView.name) {
					idx++
				}
			}
			component.componentView.idx = componentsLength
			if (idx > 0) {
				component.componentView.defaultLable = `${component.componentView.defaultLable}(${idx})`
			}

			if (sort > 0) {
				list.splice(sort, 0, component.componentView)
			} else if (sort == 0) {
				list.unshift(component.componentView)
			}
			self.selected = list.indexOf(component.componentView)
			drag.$emit("selectComponent", component.componentView)
		},
		getComponents() {
			let count = 0
			for (let i = 0, l = this.domArr.length; i < l; i++) {
				count++
				if (Object.prototype.toString.call(this.domArr[i].domArr).slice(8, -1) === "Array") {
					for (let m = 0, n = this.domArr[i].domArr.length; m < n; m++) {
						count++
					}
				}
			}
			return count
		},
		isInCanvas(obj) {
			let self = this
			return obj.clientX >= self.left && obj.clientY >= (self.top - self.scrollTop) && obj.clientX <= (self.left + self.width) && obj.clientY <= (self.top + self.height - self.scrollTop)
		}
  },
  watch: {
    list: {
      handler (val) {
        this.components = val
      },
      deep: true
    },
    scrollTop(){
    	this.scrollTop =  this.frameScrollTop + this.canvasScrollTop
    }
  },
  created() {
		let self = this
		//需要在这里获取到已选择的组件
		this.components = this.list

		drag.$on("moveInCanvas", function (obj) {
			if (self.isInCanvas(obj)) {
				let topInCanvas = obj.clientY - self.top + self.scrollTop

				if (self.domArr.length <= 0) {
					self.inCanvas = 0;
					self.tabIndex = null;
				} else if (self.domArr.length === 1) {

					if (obj.componentName == 'tablefield') {
						if (topInCanvas <= self.domArr[0].middle) {
							self.inCanvas = 0
							self.tabIndex = null
						} else if (topInCanvas > self.domArr[0].middle) {
							self.inCanvas = 1
							self.tabIndex = null
						}
					} else {
						if (topInCanvas <= self.domArr[0].middle_top) {
							self.inCanvas = 0
							self.tabIndex = null
						} else if (topInCanvas > self.domArr[0].middle_lower) {
							self.inCanvas = 1
							self.tabIndex = null
						} else if (topInCanvas <= self.domArr[0].middle_lower && topInCanvas > self.domArr[0].middle_top) {
							// 处于table中
							let item = self.domArr[0]
							self.inCanvas = null
							self.tabIndex = 0
							if (item.domArr.length <= 0 || topInCanvas <= item.domArr[0].middle_top) {
								self.components[0].InTableCanvas = 0
							} else if (topInCanvas > item.domArr[item.domArr.length - 1].middle_lower) {
								self.components[0].InTableCanvas = item.domArr.length
							} else if (item.domArr.length > 1) {
								for (let m = 0, n = item.domArr.length - 1; m < n; m++) {
									if (topInCanvas > item.domArr[m].middle_lower && topInCanvas <= item.domArr[m + 1].middle_top) {
										self.components[0].InTableCanvas = m + 1
									}
								}
							}
						}
					}
				} else {
					// console.log('topInCanvas=clientY=top=self.domArr[0].middle',topInCanvas, obj.clientY, self.top, self.domArr[0].middle)

					if (obj.componentName == 'tablefield') {
						if (topInCanvas <= self.domArr[0].middle) {
							self.inCanvas = 0
							self.tabIndex = null
						} else if (topInCanvas > self.domArr[self.domArr.length - 1].middle) {
							self.inCanvas = self.domArr.length
							self.tabIndex = null
						} else {
							for (let i = 0, l = self.domArr.length; i < l - 1; i++) {
								let item = self.domArr[i]
								let nextItem = self.domArr[i + 1]
								if (topInCanvas > item.middle && topInCanvas <= nextItem.middle) {
									self.inCanvas = i + 1
									self.tabIndex = null
									self.components[i].InTableCanvas = null
								}
							}
						}
					} else {
						if (topInCanvas <= self.domArr[0].middle_top) {
							self.inCanvas = 0
							self.tabIndex = null
						} else if (topInCanvas > self.domArr[self.domArr.length - 1].middle_lower) {
							self.inCanvas = self.domArr.length
							self.tabIndex = null
						} else {
							for (let i = 0, l = self.domArr.length; i < l - 1; i++) {
								let item = self.domArr[i]
								let nextItem = self.domArr[i + 1]
								// 处于table中
								if (topInCanvas > item.middle_top && topInCanvas <= item.middle_lower) {
									self.inCanvas = null
									self.tabIndex = i
									if (item.domArr.length <= 0 || topInCanvas <= item.domArr[0].middle_top) {
										self.components[i].InTableCanvas = 0
									} else if (topInCanvas > item.domArr[item.domArr.length - 1].middle_lower) {
										self.components[i].InTableCanvas = item.domArr.length
									} else {
										for (let m = 0, n = item.domArr.length - 1; m < n; m++) {
											if (topInCanvas > item.domArr[m].middle_lower && topInCanvas <= item.domArr[m + 1].middle_top) {
												self.components[i].InTableCanvas = m + 1
											}
										}
									}
								} else if (topInCanvas > item.middle_lower && topInCanvas <= nextItem.middle_top) {
									self.inCanvas = i + 1
									self.tabIndex = null
									self.components[i].InTableCanvas = null
								}

							}

							//当最后一个组件是明table时
							if (topInCanvas <= self.domArr[self.domArr.length - 1].middle_lower && topInCanvas > self.domArr[self.domArr.length - 1].middle_top) {
								let item = self.domArr[self.domArr.length - 1]
								self.inCanvas = null
								self.tabIndex = self.domArr.length - 1
								if (item.domArr.length <= 0 || topInCanvas <= item.domArr[0].middle_top) {
									self.components[self.domArr.length - 1].InTableCanvas = 0
								} else if (topInCanvas > item.domArr[item.domArr.length - 1].middle_lower) {
									self.components[self.domArr.length - 1].InTableCanvas = item.domArr.length
								} else {
									for (let m = 0, n = item.domArr.length - 1; m < n; m++) {
										if (topInCanvas > item.domArr[m].middle_lower && topInCanvas <= item.domArr[m + 1].middle_top) {
											self.components[m].InTableCanvas = m + 1
										}
									}
								}
							}
						}
					}
				}

				// table重置
				if (self.tabIndex != null) {
					self.inCanvas = null
				} else {
					self.inCanvas = self.inCanvas || 0
				}

				if (self.inCanvas != null) {
					for (let i = 0, l = self.components.length; i < l; i++) {
						self.components[i].InTableCanvas = null
					}
				}
			} else {
				// 重置
				self.inCanvas = null
				self.tabIndex = null
				for (let i = 0, l = self.components.length; i < l; i++) {
					self.components[i].InTableCanvas = null
				}
			}
		})

		drag.$on("moveEnd", function (obj) {
			// console.log('inCanvas=tabIndex=parNodeIndex->', self.inCanvas, self.tabIndex, self.parNodeIndex)

			if (self.inCanvas != null) {
				if (self.isDrag) {
					self.dragIndex = self.dragIndex >> 0;
					
					if (self.parNodeIndex != null) {
						console.log('1、从table里面拖动外围')
						let dragItem = self.components[self.parNodeIndex].components[self.dragIndex]
						if (!dragItem) return
						self.components[self.parNodeIndex].components.splice(self.dragIndex, 1)
						self.components.splice(self.inCanvas, 0, dragItem)
						self.selected = self.components.indexOf(dragItem)
					} else {
					console.log('2、外围互相拖动	')
						if (self.dragIndex != self.inCanvas - 1 && self.dragIndex != self.inCanvas) {
							let dragItem = self.components[self.dragIndex]
							self.components.splice(self.dragIndex, 1)
							self.components.splice(self.inCanvas, 0, dragItem)
							self.selected = self.components.indexOf(dragItem)
						}
					}

					self.dragIndex = null
					self.isDrag = false
				} else {
					console.log('添加新组件')
					self.addComponents(obj, self.components, self.inCanvas)
				}
				self.inCanvas = null

			} else if (self.tabIndex != null) {
				if (self.isDrag) {
					let inTabIndex = self.components[self.tabIndex].InTableCanvas >> 0
					self.dragIndex = self.dragIndex >> 0

					if (self.parNodeIndex == null) {
						console.log('3、从外围拖到table里面')
						let dragItem = self.components[self.dragIndex]
						if (!dragItem) return
						self.components[self.tabIndex].components.splice(inTabIndex, 0, dragItem)
						self.components[self.tabIndex].selected = inTabIndex
						self.components.splice(self.dragIndex, 1)
						self.selected = null
					} else {
						console.log('4、table里面互相拖动')
						if ((self.parNodeIndex == self.tabIndex && self.dragIndex != inTabIndex && self.dragIndex != inTabIndex - 1) || self.parNodeIndex != self.tabIndex) {
							let dragItem = self.components[self.parNodeIndex].components[self.dragIndex]
							self.components[self.parNodeIndex].components.splice(self.dragIndex, 1)
							if (inTabIndex > 0) {
								self.components[self.tabIndex].components.splice(inTabIndex, 0, dragItem)
							} else if (inTabIndex == 0) {
								self.components[self.tabIndex].components.unshift(dragItem)
							}
							self.selected = null
							self.components[self.tabIndex].selected = inTabIndex
						}
					} 

					self.dragIndex = null
					self.isDrag = false
				} else {
					console.log('添加到table')
					let list = self.components[self.tabIndex].components
					let sort = self.components[self.tabIndex].InTableCanvas
					self.addComponents(obj, list, sort)
				}
				
				for (let i = 0, l = self.components.length; i < l; i++) {
					self.components[i].InTableCanvas = null
				}
			}
			
			if (self.components.length <= 0) {
				self.isEmpty = true
			} else {
				self.isEmpty = false
			}

			document.querySelectorAll('.wf-component').forEach(it=>{
				it.classList.remove('draging')
			})

			drag.$emit("dragend", obj)
		})

		drag.$on("changeComponent", function (obj) {
			for (let i = 0; i < self.components.length; i++) {
				let item = self.components[i]
				if (item.idx == obj.idx) {
					self.components.splice(i, 1, obj)
				} else if (item.componentName == 'tablefield') {
					for (let m = 0, n = item.components.length; m < n; m++) {
						if (obj.idx == item.components[m].idx) {
							self.components[i].components.splice(m, 1, obj)
						}
					}
				}
			}
		})

		drag.$emit('save', self.components)
  },
  mounted() {
  	if(!this.frame) this.frame = document.querySelector('.xa-frame__main')
  	if(!this.canvas) this.canvas = document.querySelector('.wf-canvas-inner')
		let dom = document.querySelector('.wf-canvas-inner')
		let actualLeft = dom.offsetLeft
		let actualTop = dom.offsetTop
		let current = dom.offsetParent
		while (current !== null) {
			actualLeft += current.offsetLeft
			actualTop += current.offsetTop
			current = current.offsetParent
		}
		this.left = actualLeft
		this.top = actualTop
		this.scrollTop = this.frame.scrollTop
		this.width = dom.offsetWidth
		this.height = dom.offsetHeight

		this.frame.onscroll = ()=>{
			this.frameScrollTop = this.frame.scrollTop
			this.scrollTop = this.canvasScrollTop + this.frameScrollTop
		}

		this.canvas.onscroll = ()=>{
			this.canvasScrollTop = this.canvas.scrollTop
			this.scrollTop = this.canvasScrollTop + this.frameScrollTop
		}
	},
  updated() {
		this.domArr = []
		let domArr = document.querySelectorAll('.wf-canvas-body>div>div>.wf-component')
		for (let i = 0, len = domArr.length; i < len; i++) {
			let obj = domArr[i]
			if (obj.className.indexOf('wf-component-tablefield') >= 0) {
				let middleDomArr = []
				let objTop = obj.offsetTop
				let middleDom = obj.querySelectorAll('.wf-component')
				for (let m = 0, n = middleDom.length; m < n; m++) {
					let item = middleDom[m]
					middleDomArr.push({
						height: item.offsetHeight,
						middle_top: (objTop + 49 + item.offsetTop + item.offsetHeight / 2) >> 0,
						middle_lower: (objTop + 49 + item.offsetTop + item.offsetHeight / 2) >> 0,
						top: item.offsetTop + objTop
					})
				}
				this.domArr.push({
					height: obj.offsetHeight,
					middle_top: (obj.offsetTop + 49) >> 0,
					middle: (obj.offsetTop + obj.offsetHeight / 2) >> 0,
					middle_lower: (obj.offsetTop + obj.offsetHeight - 54) >> 0,
					top: objTop,
					domArr: middleDomArr
				})
			} else {
				this.domArr.push({
					height: obj.offsetHeight,
					middle: (obj.offsetTop + obj.offsetHeight / 2) >> 0,
					middle_top: (obj.offsetTop + obj.offsetHeight / 2) >> 0,
					middle_lower: (obj.offsetTop + obj.offsetHeight / 2) >> 0,
					top: obj.offsetTop
				})
			}
	  }
	}
}
</script>
<style lang="scss" scoped>

.wf {

	&-canvas {
    position:absolute;
  	left:50%;
    top:0;
    bottom:0;
    margin-left:-245px;
    height:670px;
    width:490px;
    background-size:100%;
    background-image: url('../../../assets/images/workflow/phone.png');
    background-position:top center;
    background-repeat:no-repeat;
   

    &-inner {
      position: absolute;
      overflow-x: hidden;
      overflow-y: auto;
      left: 99px;
	    top: 73px;
	    bottom: 0;
	    width: 293px;
	    max-height: 482px;
      background-color: #f0eff4;
    }

    &-body {
      overflow: hidden;
      padding-bottom: 30px;

      &.empty {
      	min-height: 300px;
        background-image: url('../../../assets/images/workflow/drag.png');
        background-position:center;
        background-repeat: no-repeat;
    		position: relative;
      }
    }
  }

  &-dragging {

  	&-mark {
  		background: #ff4500;
		  height: 2px;
		  overflow: visible;
		  position: relative;
		  margin: 3px 0;

		  &::before,
		  &::after {
		  	position: absolute;
			  content: '';
			  display: block;
			  height: 0;
			  width: 0;
			  border-style: solid;
			  border-color: transparent;
			  border-width: 4px 8px;
			  top: -3px;
		  }

		  &::before {
		  	left: 0px;
  			border-left-color: #ff4500;
		  }

		  &::after {
		  	right: 0px;
  			border-right-color: #ff4500;
		  }
  	}
  }

  &-remove {
  	display: none;
	  position: absolute;
	  right: 0;
	  top: 0;
	  z-index: 10;
	  width:20px;
	  height:20px;
	  text-align:center;
	  line-height:20px;
	  font-size: 18px;
	  color: #fff;
	  cursor: pointer;
	  background: $color-primary;
	}

	&-overlay {
	  position: absolute;
	  top: 0;
	  right: 0;
	  bottom: 0;
	  left: 0;
	  z-index: 9;
	  border: 1px dashed transparent;
	  cursor: move;
	}

	&-component {
  	margin: 10px 0;
  	overflow: hidden;
  	position: relative;

  	&.draging {
		  pointer-events: none;
		  -webkit-filter: blur(1px);
		  filter: blur(1px);
		
			& > .wf-overlay {
			  border-width: 1px;
			  border-color: #666;
			  border-style: dashed;
			  background: #ccc;
			  opacity: 0.8;
			}
		}

		&-group {
			.wf-component {
				margin:8px 0;
			}
		}

		&-textarea {
			.wf-componentview-placeholder {
				display: inline-block;
				height: 50px;
			}
		}

		&-daterange {
			.wf-componentview-border:first-child {
				border-bottom: 1px solid $color-border-gray;
			}
		  
		}

		&-textnote {
			.wf-componentview-placeholder {
				justify-content:flex-start;
				padding-bottom:30px;
			}
		}

		&-phone {
			.wf-componentview-border {

				&:not(:first-child) {

					.wf-componentview-label {
						border-right:1px solid $color-border-gray;

						i {
							margin-left:4px;
						}
					}

					.wf-componentview-placeholder {
						padding-left:20px;
						justify-content:flex-start;
					}
				}
			}
		}

		&-rate {

			.wf-componentview-border:not(:first-child) {
				padding-top:0;
			}
		}

  }

	&-componentview {
	  position: relative;
	  background-color:#f0eff4;

		&-border {
		  background: #fff;
		  padding: 10px;
		  white-space: nowrap;
		  @include flex-row();
		}

		&-label {
		  display: inline-block;
		  width: 5.5em;
		  padding: 0 8px 0 0;
		  font-size:14px;
		  color: #222;
		  white-space:nowrap;

		  b {
		  	color:$color-danger;
		  }
		}

		&-placeholder {
		  color: #ccc;
		  font-size: 14px;
		  @include flex-row();
		  align-items:center;
		  flex:3;
		  text-align: right;
		  justify-content: flex-end;

			i {
				font-size:18px;
				margin-top:-3px;
			}

		  .icon-camera {
			  font-size: 24px;
			  color: #858e99;
			}
		}

		&-title {
			width: 100%;
			padding: 10px;
    	background: #fff;
		}

		&-tips {
			padding:3px 10px;
			font-size:12px;
		}

		&-rate {
			display:flex;
  		flex-direction:column;
			color:$color-999;
			flex:1;
			text-align:center;
			font-size:12px;

  		&.rate10 {
  			i {
  				font-size:24px;
  				color:$color-border-gray;
  			}
  		}

  		&.rate20 {
  			i {
  				font-size:32px;
  				color:$color-border-gray;
  			}
  		}
		}

		&-area {
		  min-height: 90px;
		  background: #ddeff3;
		  position: relative;

			& .emptytip {
			  position: absolute;
			  left: 50%;
			  top: 50%;
			  white-space: nowrap;
			  color: #ccc;
			  font-size:12px;
			  transform: translate(-50%, -50%);
			}
		}

		&-adddetail {
		  color: #008cee;
		  text-align: center;
		  padding: 10px;
		  background: #f7f9ff;
		  border-bottom: 1px solid rgba(200,200,202,0.5);
		  vertical-align: middle;
		}
	
	}
}	

html.wf-cursor-move,
html.wf-cursor-move * {
  cursor: move !important;
}
html:not(.wf-cursor-move) .wf-component.hover > .wf-remove {
  display: block;
}
html:not(.wf-cursor-move) .wf-component.hover > .wf-overlay {
  border-color: $color-primary;
}
html:not(.wf-cursor-move) .wf-component.active > .wf-overlay {
  border-width: 1px;
  border-style: solid;
  border-color: $color-primary;
}
</style>