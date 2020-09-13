import React from 'react';
import { Link } from 'react-router-dom';
import _ from 'lodash';
import './Pagination.css';
// https://ing-yeo.net/2019/08/react-beginner-3/

interface PaginationProps {
	pages: number;
	currentPage: number;
	onPageChange: (page: number) => void;
}
const Pagination: React.FunctionComponent<PaginationProps> = (props) => {
	if (props.pages === 1) return null; // 1페이지 뿐이라면 페이지 수를 보여주지 않음

	const pages = _.range(1, props.pages + 1); // 마지막 페이지에 보여줄 컨텐츠를 위해 +1, https://lodash.com/docs/#range 참고

	return (
		<nav>
			{/* VSCode 입력: nav>ul.pagination>li.page-item>a.page-link */}
			<ul className="pagination">
				{pages.map((page: number) => (
					<li
						key={page}
						className={
							page === props.currentPage ? 'page-item active' : 'page-item'
						}
						style={{ cursor: 'pointer' }}
					>
						{/* bootstrap을 이용하여 현재 페이지를 시각적으로 표시*/}
						<button
							className="page-link"
							onClick={() => props.onPageChange(page)}
							style={
								page === props.currentPage
									? { background: 'pink', color: 'black' }
									: { color: 'black' }
							}
						>
							{page}
						</button>{' '}
						{/* 페이지 번호 클릭 이벤트 처리기 지정 */}
					</li>
				))}
			</ul>
		</nav>
	);
};

export default Pagination;
