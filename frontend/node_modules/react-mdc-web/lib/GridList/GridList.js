'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var ROOT = 'mdc-grid-list';
var TILES = ROOT + '__tiles';
var COMPACT = ROOT + '--tile-gutter-1';
var HEADER_CAPTION = ROOT + '--header-caption';
var RATIO = ROOT + '--tile-aspect';
var TWOLINE = ROOT + '--twoline-caption';
var ICON_ALIGN_START = ROOT + '--with-icon-align-start';
var ICON_ALIGN_END = ROOT + '--with-icon-align-end';

var GridList = function (_PureComponent) {
  _inherits(GridList, _PureComponent);

  function GridList(props) {
    _classCallCheck(this, GridList);

    var _this = _possibleConstructorReturn(this, (GridList.__proto__ || Object.getPrototypeOf(GridList)).call(this, props));

    _this.state = {};
    _this.handleResize = _this.handleResize.bind(_this);
    return _this;
  }

  _createClass(GridList, [{
    key: 'componentDidMount',
    value: function componentDidMount() {
      window.addEventListener('resize', this.handleResize);
    }
  }, {
    key: 'componentWillUnmount',
    value: function componentWillUnmount() {
      window.removeEventListener('resize', this.handleResize);
    }
  }, {
    key: 'handleResize',
    value: function handleResize() {
      this.forceUpdate();
    }
  }, {
    key: 'render',
    value: function render() {
      var _this2 = this,
          _classnames;

      var _props = this.props,
          captionIconAlign = _props.captionIconAlign,
          children = _props.children,
          className = _props.className,
          compact = _props.compact,
          headerCaption = _props.headerCaption,
          ratio = _props.ratio,
          twoLineCaption = _props.twoLineCaption,
          otherProps = _objectWithoutProperties(_props, ['captionIconAlign', 'children', 'className', 'compact', 'headerCaption', 'ratio', 'twoLineCaption']);

      var childs = _react.Children.map(children, function (child, index) {
        if (index === 0) {
          return (0, _react.cloneElement)(child, {
            onWidthChange: function onWidthChange(width) {
              _this2.setState({ itemWidth: width });
            }
          });
        }
        return child;
      });

      var _state = this.state,
          gridWidth = _state.gridWidth,
          itemWidth = _state.itemWidth;

      var tilesWidth = itemWidth * Math.floor(gridWidth / itemWidth);

      return _react2.default.createElement(
        'div',
        _extends({
          className: (0, _classnames3.default)(ROOT, (_classnames = {}, _defineProperty(_classnames, COMPACT, compact), _defineProperty(_classnames, HEADER_CAPTION, headerCaption), _defineProperty(_classnames, TWOLINE, twoLineCaption), _defineProperty(_classnames, RATIO + '-' + ratio, ratio), _defineProperty(_classnames, ICON_ALIGN_START, captionIconAlign === 'start'), _defineProperty(_classnames, ICON_ALIGN_END, captionIconAlign === 'end'), _classnames), className),
          ref: function ref(native) {
            _this2.setState({ gridWidth: native && native.offsetWidth });
          }
        }, otherProps),
        _react2.default.createElement(
          'ul',
          {
            className: TILES,
            style: { width: tilesWidth + 'px' }
          },
          childs
        )
      );
    }
  }]);

  return GridList;
}(_react.PureComponent);

GridList.propTypes = {
  captionIconAlign: _propTypes2.default.oneOf(['start', 'end']),
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  compact: _propTypes2.default.bool,
  headerCaption: _propTypes2.default.bool,
  ratio: _propTypes2.default.oneOf(['1x1', '16x9', '2x3', '3x2', '4x3', '3x4']),
  twoLineCaption: _propTypes2.default.bool
};
exports.default = GridList;