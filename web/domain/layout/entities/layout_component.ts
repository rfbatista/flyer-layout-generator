export type LayoutComponentProps = {
  id: number
  elements: Array<{
    id: number
    xi: number
    xii: number
    yi: number
    yii: number
    inner_xii: number
    inner_yii: number
    layer_id: string
    width: number
    height: number
    kind: string
    name: string
    is_group: boolean
    group_id: number
    level: number
    photoshop_id: number
    image: string
    component_id: number
    InnerContainer: {
      UpperLeft: {
        X: number
        Y: number
      }
      DownRight: {
        X: number
        Y: number
      }
    }
    OuterContainer: {
      UpperLeft: {
        X: number
        Y: number
      }
      DownRight: {
        X: number
        Y: number
      }
    }
  }>
  width: number
  height: number
  type: string
  xii: number
  yii: number
  bbox_xii: number
  bbox_yii: number
  left_gap: {
    X: number
    Y: number
  }
  right_gap: {
    X: number
    Y: number
  }
  up_gap: {
    X: number
    Y: number
  }
  down_gap: {
    X: number
    Y: number
  }
  InnerContainer: {
    UpperLeft: {
      X: number
      Y: number
    }
    DownRight: {
      X: number
      Y: number
    }
  }
  OuterContainer: {
    UpperLeft: {
      X: number
      Y: number
    }
    DownRight: {
      X: number
      Y: number
    }
  }
  GridContainer: {
    UpperLeft: {
      X: number
      Y: number
    }
    DownRight: {
      X: number
      Y: number
    }
  }
  Priority: number
  Positions: any
  Pivot: {
    X: number
    Y: number
  }
}

export class LayoutComponent {

}
