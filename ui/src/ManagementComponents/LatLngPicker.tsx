import React, { Component } from 'react';
import { locationElem } from '../ElemInterface/ElementsInterface';
// 참고 : https://apis.map.kakao.com/web/sample/addMapClickEventWithMarker/
// 참고 : https://chaewonkong.github.io/posts/react-kakao-maps.html

declare global {
	interface Window {
		kakao: any;
	}
}
interface LarLngPickerProps {
	handleLarLngChange: (location: locationElem) => void;
}

class LarLngPicker extends Component<LarLngPickerProps, {}> {
	componentDidMount() {
		var mapContainer = document.getElementById('map'); // 지도를 표시할 div
		var mapOption = {
			center: new window.kakao.maps.LatLng(
				37.49575158172499,
				126.95633291769067
			), // 지도의 중심좌표
			level: 3, // 지도의 확대 레벨
		};

		// 지도를 생성합니다
		var map = new window.kakao.maps.Map(mapContainer, mapOption);

		// 지도를 클릭한 위치에 표출할 마커입니다
		var marker = new window.kakao.maps.Marker({
			// 지도 중심좌표에 마커를 생성합니다
			position: map.getCenter(),
		});

		// 지도에 마커를 표시합니다
		marker.setMap(map);

		// 지도에 클릭 이벤트를 등록합니다
		// 지도를 클릭하면 마지막 파라미터로 넘어온 함수를 호출합니다
		window.kakao.maps.event.addListener(map, 'click', (mouseEvent: any) => {
			// 클릭한 위도, 경도 정보를 가져옵니다
			var latlng = mouseEvent.latLng;

			// 마커 위치를 클릭한 위치로 옮깁니다
			marker.setPosition(latlng);

			// re layout map (bug)
			map.relayout();

			// loc : save latitude and longitude which user pick
			var loc = { lat: latlng.getLat(), lng: latlng.getLng() };

			// send loc to parent component( RegisterNode )
			this.props.handleLarLngChange(loc);

			var message = 'latitude : ' + latlng.getLat() + ', ';
			message += 'longitude : ' + latlng.getLng() + '\n';

			var resultDiv: any = document.getElementById('clickLatlng');
			resultDiv.innerHTML = message;
		});
	}

	render() {
		return (
			<div>
				<div
					id="map"
					style={{ position: 'relative', width: '100%', height: '350px' }}
				></div>
				<p>
					<em>Click map!</em>
				</p>
				<div id="clickLatlng"></div>
				<br />
			</div>
		);
	}
}

export default LarLngPicker;
