import React, { Component } from 'react';
import './NodeMap.css';
import { NODE_URL } from '../defineUrl';

// 참고 : https://apis.map.kakao.com/web/sample/multipleMarkerEvent/

import {
	nodeListElem,
	sensorListElem,
	nodeHealthCheckElem,
} from '../ElemInterface/ElementsInterface';
import MapNodeTable from './Table/MapNodeTable';
declare global {
	interface Window {
		kakao: any;
	}
}

interface NodeMapProps {
	nodeState: Array<nodeHealthCheckElem>;
}

interface NodeMapState {
	nodeList: Array<nodeListElem>;
	map: any;
	left: number;
	right: number;
	up: number;
	down: number;
}

class NodeMap extends Component<NodeMapProps, NodeMapState> {
	state: NodeMapState = {
		nodeList: [],
		map: {},
		left: 126.93120847993194,
		right: 126.9814068917757,
		up: 37.504736714448086,
		down: 37.48669801512536,
	};
	componentDidMount = () => {
		var mapContainer = document.getElementById('node_map'); // 지도를 표시할 div
		var mapOption = {
			center: new window.kakao.maps.LatLng(
				37.49575158172499,
				126.95633291769067
			), // 지도의 중심좌표
			level: 5, // 지도의 확대 레벨
		};

		// 지도를 생성합니다
		var map = new window.kakao.maps.Map(mapContainer, mapOption);
		this.setState({ map: map });

		this.getnodeList(
			this.state.left,
			this.state.right,
			this.state.up,
			this.state.down,
		);

		// 드래그가 끝날 때 or 확대 수준이 변경되면
		window.kakao.maps.event.addListener(map, 'bounds_changed', () => {
			// 지도의 현재 영역을 얻어옵니다
			var bounds = map.getBounds();

			// 영역의 남서쪽 좌표를 얻어옵니다
			var swLatLng = bounds.getSouthWest();

			// 영역의 북동쪽 좌표를 얻어옵니다
			var neLatLng = bounds.getNorthEast();

			this.setState({
				left: swLatLng.getLng(), // left
				right: neLatLng.getLng(), // right
				up: neLatLng.getLat(), // up
				down: swLatLng.getLat(), // down
			});

			this.getnodeList(
				swLatLng.getLng(), // left
				neLatLng.getLng(), // right
				neLatLng.getLat(), // up
				swLatLng.getLat() // down
			);
		});
	};
	// Get node list from backend
	getnodeList(left: number, right: number, up: number, down: number) {
		var url =
			NODE_URL +
			'?left=' +
			left +
			'&right=' +
			right +
			'&up=' +
			up +
			'&down=' +
			down;
		fetch(url)
			.then((res) => res.json())
			.then((data) => this.setState({ nodeList: data }))
			.catch((error) => console.error('Error:', error));
	}

	// Make Marker
	displayMarker(position: any, map: any) {
		// 마커를 생성합니다
		var marker = new window.kakao.maps.Marker({
			map: map, // 마커를 표시할 지도
			position: position.latlng, // 마커의 위치
		});

		// 마커에 표시할 custom overlay를 생성합니다
		var customOverlay = new window.kakao.maps.CustomOverlay({
			position: marker.getPosition(),
		});

		var content = document.createElement('div');

		var title = document.createElement('div');
		title.innerHTML = position.title;
		title.className = 'title';

		var closeBtn = document.createElement('button');

		closeBtn.onclick = function () {
			customOverlay.setMap(null);
		};
		closeBtn.className = 'close';
		closeBtn.type = 'button';
		title.appendChild(closeBtn);

		var body = document.createElement('div');

		for (var i = 0; i < 3; i++) {
			var bodyElem = document.createElement('div');
			bodyElem.appendChild(document.createTextNode(position.content[i]));
			body.insertAdjacentElement('beforeend', bodyElem);
		}

		body.className = 'body';

		var info = document.createElement('div');
		info.className = 'info';

		info.insertAdjacentElement('afterbegin', title);

		var wrap = document.createElement('div');
		wrap.className = 'wrap';
		wrap.insertAdjacentElement('afterbegin', info);
		title.insertAdjacentElement('afterend', body);
		content.insertAdjacentElement('afterbegin', wrap);

		customOverlay.setContent(content);

		window.kakao.maps.event.addListener(marker, 'click', function () {
			customOverlay.setMap(map);
		});
	}

	render() {
		var positions = this.state.nodeList.map((node: nodeListElem) => {
			return {
				title: node.name,
				content: [
					'sink : ' + node.sink_id,
					'id : ' + node.id,
					'sensor : ' +
						node.sensors.map((sensor: sensorListElem) => sensor.name),
				],
				latlng: new window.kakao.maps.LatLng(node.lat, node.lng),
			};
		});

		for (var i = 0; i < positions.length; i++) {
			this.displayMarker(positions[i], this.state.map);
		}

		return (
			<div>
				<div>
					<div id="node_map" style={{ width: '100%', height: '500px' }}></div>
					<MapNodeTable
						nodeState={this.props.nodeState}
						nodeList={this.state.nodeList}
					></MapNodeTable>
				</div>
			</div>
		);
	}
}

export default NodeMap;
