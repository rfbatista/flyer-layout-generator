type Props = {
  page: number;
  limit: number;
  total: number;
};

export function Pagination(props: Props) {
  if (props.total <= props.limit) {
    return
  }
  return (
    <nav aria-label="pagination">
      <ul data-type="pagination">
        <li>
          <button data-type="icon">
            <span data-type="chevron-left" />
            Previous
          </button>
        </li>
        <li>
          <button data-type="icon">
            Next
            <span data-type="chevron-right" />
          </button>
        </li>
      </ul>
    </nav>
  );
}
